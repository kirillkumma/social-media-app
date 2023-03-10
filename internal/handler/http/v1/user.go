package v1

import (
	"github.com/gofiber/fiber/v2"
	"social-media-app/internal/usecase"
)

type UserHandler struct {
	createUserInteractor usecase.CreateUserInteractor
}

func (h *UserHandler) create(ctx *fiber.Ctx) error {
	var p usecase.CreateUserInputDTO
	if err := ctx.BodyParser(&p); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(newErrResp(err))
	}

	output, err := h.createUserInteractor.Execute(ctx.Context(), p)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(newResp(output))
}

func (h *UserHandler) Register(r fiber.Router) {
	r.Post("", h.create)
}

func NewUserHandler(createUserInteractor usecase.CreateUserInteractor) *UserHandler {
	return &UserHandler{createUserInteractor}
}
