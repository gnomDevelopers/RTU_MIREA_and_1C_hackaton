package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"server/internal/log"
	"strconv"
)

// CreateUniversity
// @Tags university
// @Summary      Create university
// @Accept       json
// @Produce      json
// @Param data body entities.CreateUniversityRequest true "university data"
// @Success 200 {object} entities.CreateUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /university [post]
func (h *Handler) CreateUniversity(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль админа
	var university entities.University
	err := c.BodyParser(&university)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.UniversityService.Create")
	id, err := h.services.UniversityService.Create(c.Context(), university.Name)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

// GetAllUniversities
// @Tags university
// @Summary      Get all universities
// @Accept       json
// @Produce      json
// @Success 200 {object} []entities.University
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /university/all [get]
func (h *Handler) GetAllUniversities(c *fiber.Ctx) error {
	h.logger.Debug().Msg("call h.services.UniversityService.GetAll")
	universities, err := h.services.UniversityService.GetAll(c.Context())
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(universities)
}

// GetByNameUniversity
// @Tags university
// @Summary      Get university by name
// @Accept       json
// @Produce      json
// @Param name path string true "university name"
// @Success 200 {object} entities.University
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /university/name/{name} [get]
func (h *Handler) GetByNameUniversity(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль админа
	name := c.Params("name")

	h.logger.Debug().Msg("call h.services.UniversityService.GetByName")
	university, err := h.services.UniversityService.GetByName(c.Context(), name)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(university)
}

// UpdateUniversity
// @Tags university
// @Summary      Update university
// @Accept       json
// @Produce      json
// @Param data body entities.University true "university data"
// @Success 200 {object} entities.UpdateDeleteUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /university [put]
func (h *Handler) UpdateUniversity(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль админа
	var university entities.University
	err := c.BodyParser(&university)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.UniversityService.Update")
	err = h.services.UniversityService.Update(c.Context(), university.Id, university.Name)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("university with id=%v updated successfully", university.Id)})
}

// DeleteUniversity
// @Tags university
// @Summary      Delete university
// @Accept       json
// @Produce      json
// @Param id path string true "university id"
// @Success 200 {object} entities.UpdateDeleteUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /university/{id} [delete]
func (h *Handler) DeleteUniversity(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль админа
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid id in path")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id in path"})
	}

	h.logger.Debug().Msg("call h.services.UniversityService.Delete")
	err = h.services.UniversityService.Delete(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("university with id=%v deleted successfully", id)})
}
