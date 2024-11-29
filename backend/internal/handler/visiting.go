package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"strconv"
)

// CheckIn
// @Tags visiting
// @Summary      Record a check-in for a student
// @Accept       json
// @Produce      json
// @Param data body entities.CheckInRequest true "Check-in data"
// @Success 200 {object} map[string]interface{} "status ok"
// @Failure 400 {object} map[string]interface{} "status bad request"
// @Failure 500 {object} map[string]interface{} "status internal server error"
// @Router       auth/class/visiting/check-in [post]
// @Security ApiKeyAuth
func (h *Handler) CheckIn(c *fiber.Ctx) error {
	var req entities.CheckInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.services.VisitingService.Update(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}

// GetClassVisiting
// @Tags visiting
// @Summary      Get visiting information for a specific class
// @Accept       json
// @Produce      json
// @Param id path int true "Class ID"
// @Success 200 {object} []entities.VisitingInfo
// @Failure 400 {object} fiber.Map{"error": string}
// @Failure 500 {object} fiber.Map{"error": string}
// @Router       /auth/class/{id}/visiting [get]
// @Security ApiKeyAuth
func (h *Handler) GetClassVisiting(c *fiber.Ctx) error {
	classIDStr := c.Params("id")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.services.VisitingService.GetClassVisiting(c.Context(), classID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
