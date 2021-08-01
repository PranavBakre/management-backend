package user

import (
	"github.com/PranavBakre/management-backend/config"
	"github.com/PranavBakre/management-backend/models"
	"github.com/m4rw3r/uuid"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type APIService interface {
	Create(ctx *fiber.Ctx) error
	ReadAll(ctx *fiber.Ctx) error
	Read(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type Service struct {
	DB     *gorm.DB
	Config config.Config
}

/*
Create will add a new user to DB and return the created user
*/
func (svc *Service) Create(ctx *fiber.Ctx) error {
	// Read user data from request body
	var user models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		log.Println(err)
		return err
	}

	// Add user to DB
	result := svc.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return ctx.Status(500).SendString("User could not be added to DB")
	}

	return ctx.Status(201).JSON(user)
}

/*
ReadAll returns all users in the DB
*/
func (svc *Service) ReadAll(ctx *fiber.Ctx) error {
	var users []models.User

	// Read all users from DB
	result := svc.DB.Find(&users)
	if result.Error != nil {
		return result.Error
	}

	return ctx.JSON(users)
}

/*
Read returns a single user based on ID passed in params
*/
func (svc *Service) Read(ctx *fiber.Ctx) error {
	// Get ID from params
	id, err := uuid.FromString(ctx.Params("id"))
	if err != nil {
		log.Println(err)
		return ctx.Status(400).SendString(err.Error())
	}

	// Fetch user from DB by ID
	var user models.User
	result := svc.DB.Where("id = ?", id.String()).Find(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return ctx.Status(404).SendString("User with ID " + id.String() + " not found")
	}

	return ctx.JSON(user)
}

/*
Update will update a user record in DB and return updated object
*/
func (svc *Service) Update(ctx *fiber.Ctx) error {
	// Read user data from request body
	var user models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		log.Println(err)
		return err
	}

	// Update user in DB
	result := svc.DB.Where("id = ?", user.ID.String()).Updates(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return ctx.Status(404).SendString("User with ID " + user.ID.String() + " not found")
	}

	return ctx.JSON(user)
}

/*
Delete a user from the DB based on ID passed in params
*/
func (svc *Service) Delete(ctx *fiber.Ctx) error {
	// Get ID from params
	id, err := uuid.FromString(ctx.Params("id"))
	if err != nil {
		log.Println(err)
		return ctx.Status(400).SendString(err.Error())
	}

	// Delete user from DB by ID
	result := svc.DB.Where("id = ?", id.String()).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return ctx.Status(404).SendString("User with ID " + id.String() + " not found")
	}

	return ctx.SendStatus(200)
}
