package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
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
// @Security ApiKeyAuth
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

// GetAllGroups
// @Tags         groups
// @Summary      Get all groups by university name
// @Description  Retrieves a list of all groups associated with the specified university name.
// @Accept       json
// @Produce      json
// @Param        university path string true "University name (URL-encoded)"
// @Success      200 {object} []entities.Group "List of groups"
// @Failure      400 {object} map[string]interface{} "Invalid university name"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/groups/university/{university} [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllGroups(c *fiber.Ctx) error {
	name := c.Params("university")
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req := &entities.GetGroupRequest{
		Name: decodedName,
	}
	res, err := h.services.GroupService.GetAll(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
