package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
)

// CreateGroup
// @Tags group
// @Summary      Create a new group
// @Accept       json
// @Produce      json
// @Param data body entities.CreateGroupRequest true "Group data"
// @Success 200 {object} entities.CreateGroupResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/group [post]
func (h *Handler) CreateGroup(c *fiber.Ctx) error {
	req := &entities.CreateGroupRequest{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.services.GroupService.Create(c.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(res)
}
