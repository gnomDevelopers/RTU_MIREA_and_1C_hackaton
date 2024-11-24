package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
)

func (h *Handler) GetFacultiesByUniversityName(c *fiber.Ctx) error {
	req := entities.GetFacultyRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res, err := h.services.FacultyService.GetAll(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
