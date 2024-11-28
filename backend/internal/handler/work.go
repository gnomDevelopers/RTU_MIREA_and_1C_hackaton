package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
)

// LoginWork
// @Tags work
// @Summary      Login in hr
// @Accept       json
// @Produce      json
// @Param data body entities.CreateUniversityRequest true "university data"
// @Success 200 {object} entities.CreateUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /work/login [post]
// @Security ApiKeyAuth
func (h *Handler) LoginWork(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("")
}

// LoginHR
// @Tags         work
// @Summary      Hr login
// @Description  Authenticates a user and sets access and refresh tokens as cookies.
// @Accept       json
// @Produce      json
// @Param        data body entities.LoginUserRequest true "User login credentials"
// @Success      200 {object} entities.LoginUserResponse
// @Failure      400 {object} entities.ErrorResponse "Invalid request payload"
// @Failure      500 {object} entities.ErrorResponse "Internal server error"
// @Router       /work/login/hr [post]
func (h *Handler) LoginHR(c *fiber.Ctx) error {
	var user entities.LoginUserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	h.logger.Debug().Msg("call h.services.WorkService.Login")
	u, err := h.services.WorkService.Login(c.Context(), &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(u)
}

// ResponseCandidate
// @Tags work
// @Summary      Response
// @Accept       json
// @Produce      json
// @Param data body entities.CreateUniversityRequest true "university data"
// @Success 200 {object} entities.CreateUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/work/response [post]
// @Security ApiKeyAuth
func (h *Handler) ResponseCandidate(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("")
}

// GetResponseCandidate
// @Tags work
// @Summary      Response
// @Accept       json
// @Produce      json
// @Param data body entities.CreateUniversityRequest true "university data"
// @Success 200 {object} entities.CreateUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/work/response [get]
// @Security ApiKeyAuth
func (h *Handler) GetResponseCandidate(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("")
}
