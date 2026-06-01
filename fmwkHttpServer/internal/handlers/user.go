package handlers

import (
	"strconv"

	model "fmwkHttpServer/internal/models"
	"fmwkHttpServer/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(

	service *service.UserService,

) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
func (h *UserHandler) GetUsers(
	c *fiber.Ctx,
) error {

	users, err := h.service.GetUsers()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}

func (h *UserHandler) GetOneUser(
	c *fiber.Ctx,

) error {
	id, err := strconv.Atoi(
		c.Params("id"),
	)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid ID",
			},
		)
	}
	user, err := h.service.GetOneUser(id)

	if err != nil {
		return c.Status(404).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(user)
}

func (h *UserHandler) AddUser(
	c *fiber.Ctx,
) error {
	var req model.User
	// body parser basically acts as zod validation for the incoming payload and validates it against our user model
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	user, err := h.service.AddUser(req)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unexpected Server Error",
		})

	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
