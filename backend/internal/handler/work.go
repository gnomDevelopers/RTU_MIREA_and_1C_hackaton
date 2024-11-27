package handler

import (
	"github.com/gofiber/fiber/v2"
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
// @Router       /work [post]
// @Security ApiKeyAuth
func (h *Handler) LoginWork(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("")
}

func (h *Handler) CreateHR(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("")
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
