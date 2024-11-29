package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"strconv"
)

// CheckIn
// @Tags visiting
// @Summary Record a check-in for a student
// @Accept json
// @Produce json
// @Param data body entities.CheckInRequest true "Check-in data"
// @Success 200 {object} map[string]interface{} "status ok"
// @Failure 400 {object} map[string]interface{} "status bad request"
// @Failure 500 {object} map[string]interface{} "status internal server error"
// @Router /auth/class/visiting/check-in [post]
// @Security ApiKeyAuth
func (h *Handler) CheckIn(c *fiber.Ctx) error {
	var req entities.CheckInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	//err := h.services.VisitingService.Update(c.Context(), &req)
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	//}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}

// GetAllClassParticipants
// @Tags visiting
// @Summary      Get all participants of a specific class
// @Accept       json
// @Produce      json
// @Param id path int true "Class ID"
// @Success 200 {array} entities.ClassParticipant "List of participants"
// @Failure 400 {object} entities.ErrorResponse "Invalid class ID"
// @Failure 401 {object} entities.ErrorResponse "Unauthorized"
// @Failure 500 {object} entities.ErrorResponse "Internal server error"
// @Router /auth/class/{id}/participants [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllClassParticipants(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid class ID"})
	}

	res, err := h.services.VisitingService.Get(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// AddClassVisiting
// @Tags visiting
// @Summary      Add visiting records for a class
// @Accept       json
// @Produce      json
// @Param data body []entities.Visiting true "List of visiting records"
// @Success 200 {object} map[string]interface{} "ok"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router       /auth/class/visiting [post]
// @Security ApiKeyAuth
func (h *Handler) AddClassVisiting(c *fiber.Ctx) error {
	var visits []entities.Visiting
	if err := c.BodyParser(&visits); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.services.VisitingService.Add(c.Context(), visits)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}
