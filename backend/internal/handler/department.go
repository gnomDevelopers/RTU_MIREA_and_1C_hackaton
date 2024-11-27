package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
	"server/internal/log"
)

// GetByUniversityDepartments
// @Tags department
// @Summary      Get departments by university
// @Accept       json
// @Produce      json
// @Param university path string true "university"
// @Success 200 {object} []entities.Department
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/department/university/{university} [get]
// @Security ApiKeyAuth
func (h *Handler) GetByUniversityDepartments(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	name := c.Params("university")
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid university name")
	}

	campuses, err := h.services.DepartmentService.GetByUniversity(c.Context(), decodedName)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(campuses)
}
