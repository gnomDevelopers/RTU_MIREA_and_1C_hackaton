package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
	"server/internal/entities"
	"server/internal/log"
)

// CreateGrade
// @Tags grade
// @Summary      Create grade
// @Accept       json
// @Produce      json
// @Param data body []entities.CreateGradeRequest true "grade data"
// @Success 200 {object} []entities.CreateCampusesResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/grade [post]
// @Security ApiKeyAuth
func (h *Handler) CreateGrade(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль
	var grade entities.Grade
	err := c.BodyParser(&grade)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.CampusService.Create")
	id, err := h.services.GradeService.Create(c.Context(), &grade)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(id)
}

// GetGradesBySubject
// @Tags grade
// @Summary      Get grade by group
// @Accept       json
// @Produce      json
// @Param group path string true "group"
// @Param name path string true "name"
// @Success 200 {object} []entities.Campus
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/grade/{group}/{name} [get]
// @Security ApiKeyAuth
func (h *Handler) GetGradesBySubject(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль
	group := c.Params("group")
	decodedGroup, err := url.QueryUnescape(group)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid group")
	}
	name := c.Params("name")
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid name")
	}

	classes, err := h.services.ClassService.GetByNameAndGroup(c.Context(), decodedName, decodedGroup)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	for _, class := range *classes {
		grade, _ := h.services.GradeService.GetByUserIdAndClassId(c.Context(), class.Id, 2)
		if grade != nil {
			class.Grades = append(class.Grades, *grade)
		}
	}

	return c.Status(fiber.StatusOK).JSON(classes)
}
