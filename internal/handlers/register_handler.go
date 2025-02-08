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

func (r *RegisterHandler) HandleRegister(c *fiber.Ctx) error {
	return utils.SendSuccessResponse(c, fiber.StatusOK, "Registrado com sucesso", nil)
}
