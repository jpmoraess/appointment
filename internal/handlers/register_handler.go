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

// RegisterRequest structure for registration request
//	@Description	Object containing data for user registration
type RegisterRequest struct {
	// Name
	// @example "John"
	Name     string `json:"name"`
	// Password
	// @example "mysecr3tpwd"
	Password string `json:"password"`
}

// RegisterResponse structure for registration response
//	@Description	Response containing tenant and user information after registration
type RegisterResponse struct {
	// Tenant Identifier
	//	@example	"1"
	TenantID int64 `json:"tenantId"`
	// User Identifier
	//	@example	"1"
	UserID   int64 `json:"userId"`
}

//	@Summary		Register
//	@Description	Register new user
//	@Tags			AUTH
//	@Accept			json
//	@Produce		json
//	@Param			request	body		RegisterRequest		true	"request body"
//	@Success		200		{object}	RegisterResponse	"response data"
//	@Router			/auth/register [post]
func (r *RegisterHandler) HandleRegister(c *fiber.Ctx) error {
	var request RegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	output, err := r.register.Execute(c.Context(), &usecases.RegisterInput{
		Name:     request.Name,
		Password: request.Password,
	})
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	resp := &RegisterResponse{
		TenantID: output.TenantID,
		UserID:   output.UserID,
	}

	return utils.SendSuccessResponse(c, fiber.StatusOK, resp)
}
