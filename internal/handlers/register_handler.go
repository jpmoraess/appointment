package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/appointment-api/internal/usecases"
	"github.com/jpmoraess/appointment-api/pkg/utils"
)

type RegisterHandler struct {
	register *usecases.Register
}

func NewRegisterHandler(register *usecases.Register) *RegisterHandler {
	return &RegisterHandler{register: register}
}

type registerRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r *RegisterHandler) HandleRegister(c *fiber.Ctx) error {
	var request registerRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err := r.register.Execute(c.Context(), &usecases.RegisterInput{
		Name:     request.Name,
		Password: request.Password,
	})
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendSuccessResponse(c, fiber.StatusOK, nil)
}
