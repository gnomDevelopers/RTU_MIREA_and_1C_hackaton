package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"strconv"
)

// GetAchievementByUserID
// @Tags         achievements
// @Summary      Get achievement by user ID
// @Description  Retrieves an achievement by the specified user ID.
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} []entities.Achievement "Achievement details"
// @Failure      400 {object} map[string]interface{} "Invalid user ID"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/achievements/user/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) GetAchievementByUserID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.services.AchievementService.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// CreateAchievement
// @Tags         achievements
// @Summary      Create a new achievement
// @Description  Creates a new achievement with the provided details.
// @Accept       json
// @Produce      json
// @Param        data body entities.CreateAchievementRequest true "Achievement creation data"
// @Success      200 {object} entities.CreateAchievementResponse "Created achievement details"
// @Failure      400 {object} map[string]interface{} "Invalid request body"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/achievement [post]
// @Security ApiKeyAuth
func (h *Handler) CreateAchievement(c *fiber.Ctx) error {
	var req entities.CreateAchievementRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.services.AchievementService.Create(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// DeleteAchievement
// @Tags         achievements
// @Summary      Delete an achievement
// @Description  Deletes an achievement by the specified ID.
// @Accept       json
// @Produce      json
// @Param        id path int true "Achievement ID"
// @Success      200 {object} map[string]interface{} "Deleted achievement ID"
// @Failure      400 {object} map[string]interface{} "Invalid achievement ID"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /achievement/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) DeleteAchievement(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = h.services.AchievementService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}
