package user

import (
	"log"
	"net/http"

	"management-backend/config"
	"management-backend/models"
	"management-backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type APIService interface {
	Create(ctx *fiber.Ctx) error
	ReadAll(ctx *fiber.Ctx) error
	Read(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type Handlers struct {
	DB     *gorm.DB
	Config *config.Config
}

/*
Create will add a new user to DB and return the created user
*/
func (h *Handlers) Create(ctx *fiber.Ctx) error {
	// Read user data from request body
	var user models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		log.Println(err)
		return err
	}

	// Add user to DB
	result := h.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		log.Println("No row added")
		return ctx.Status(http.StatusInternalServerError).SendString("User could not be added to DB")
	}

	// Generate JWT token
	token, err := utils.CreateToken(h.Config, user.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"token": token, "user": user})
}

/*
Read returns a single user based on ID passed in params
*/
func (h *Handlers) Read(ctx *fiber.Ctx) error {
	// Get ID from JWT
	jwtID, err := utils.GetCurrentUserID(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Get ID from params
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Check if ID from params and JWT match
	if id != jwtID {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	// Fetch user from DB by ID
	var user models.User
	result := h.DB.Where("id = ?", id.String()).Find(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return ctx.Status(http.StatusNotFound).SendString("User with ID " + id.String() + " not found")
	}

	return ctx.JSON(user)
}

/*
Update will update a user record in DB and return updated object
*/
func (h *Handlers) Update(ctx *fiber.Ctx) error {
	// Get ID from JWT
	jwtID, err := utils.GetCurrentUserID(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Read user data from request body
	var user models.User
	err = ctx.BodyParser(&user)
	if err != nil {
		log.Println(err)
		return err
	}

	// Check if ID from params and JWT match
	if user.ID != jwtID {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	// Update user in DB
	result := h.DB.Where("id = ?", user.ID.String()).Updates(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return ctx.Status(http.StatusNotFound).SendString("User with ID " + user.ID.String() + " not found")
	}

	return ctx.JSON(user)
}

/*
Delete a user from the DB based on ID passed in params
*/
func (h *Handlers) Delete(ctx *fiber.Ctx) error {
	// Get ID from JWT
	jwtID, err := utils.GetCurrentUserID(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Get ID from params
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Check if ID from params and JWT match
	if id != jwtID {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	// Delete user from DB by ID
	result := h.DB.Where("id = ?", id.String()).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return ctx.Status(http.StatusNotFound).SendString("User with ID " + id.String() + " not found")
	}

	return ctx.SendStatus(http.StatusOK)
}
