package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) GetSchedule(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON("")
}
