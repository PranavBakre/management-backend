package auth

import (
	"time"

	"management-backend/config"
	"management-backend/models"
	"management-backend/utils/googleapis"
	"management-backend/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type APIService interface {
	Login(ctx *fiber.Ctx) error
	Verify(ctx *fiber.Ctx) error
}

type Handler struct {
	DB     *gorm.DB
	Config *config.Config
}

type LoginResponse struct {
	AccessToken        string    `json:"accessToken"`
	GoogleAccessToken  string    `json:"googleAccessToken"`
	GoogleRefreshToken string    `json:"googleRefreshToken"`
	Expiry             time.Time `json:"expiry"`
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	googleApiClient := googleapis.Get()
	var requestBody map[string]string
	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return err
	}

	googleTokenResponse, err := googleApiClient.GetToken("offline", requestBody["code"])

	if err != nil {
		return err
	}

	userinfo, err := googleApiClient.FetchSelfInfo(googleTokenResponse.AccessToken)

	user := &models.User{
		Name:     userinfo.Name,
		Email:    userinfo.Email,
		GoogleID: &userinfo.Id,
		Picture:  &userinfo.Picture,
	}

	result := h.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "google_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "email", "picture"}),
	}).Create(user)

	if result.Error != nil {
		return err
	}

	token, err := jwt.CreateToken(h.Config, user.ID)

	if err != nil {
		return err
	}

	return ctx.JSON(LoginResponse{
		AccessToken:        token,
		GoogleAccessToken:  googleTokenResponse.AccessToken,
		GoogleRefreshToken: googleTokenResponse.RefreshToken,
		Expiry:             time.Now().Add(time.Duration(googleTokenResponse.Expiry) * time.Second),
	})
}
