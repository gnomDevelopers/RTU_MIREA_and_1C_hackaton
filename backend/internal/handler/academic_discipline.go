package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"server/internal/entities"
	"server/internal/log"
	"strconv"
)

// CreateAcademicDiscipline
// @Tags discipline
// @Summary      Create discipline
// @Accept       json
// @Produce      json
// @Param data body []entities.CreateAcademicDisciplineRequest true "class data"
// @Success 200 {object} []entities.CreateAcademicDisciplineResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/discipline [post]
// @Security ApiKeyAuth
func (h *Handler) CreateAcademicDiscipline(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	var academicDisciplines []entities.AcademicDiscipline
	err := c.BodyParser(&academicDisciplines)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	var ids []int
	for _, academicDiscipline := range academicDisciplines {
		h.logger.Debug().Msg("call h.services.AcademicDisciplineService.Create")
		id, err := h.services.AcademicDisciplineService.Create(c.Context(), &academicDiscipline)
		if err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		ids = append(ids, id)
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(ids)
}

// GetAcademicDisciplinesByEducationalDirectionAndSemester
// @Tags discipline
// @Summary      Get discipline by ed dir and semester
// @Accept       json
// @Produce      json
// @Param ed_dir path string true "ed_dir"
// @Param semester path string true "semester"
// @Success 200 {object} []entities.AcademicDiscipline
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/discipline/semester/{ed_dir}/{semester} [get]
// @Security ApiKeyAuth
func (h *Handler) GetAcademicDisciplinesByEducationalDirectionAndSemester(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль админа
	edDir := c.Params("ed_dir")
	decodedEdDir, err := url.QueryUnescape(edDir)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid educational direction")
	}

	semesterStr := c.Params("semester")
	semester, err := strconv.Atoi(semesterStr)

	h.logger.Debug().Msg("call h.services.UniversityService.GetByName")
	academicDisciplines, err := h.services.AcademicDisciplineService.GetByEducationalDirectionAndSemester(c.Context(), decodedEdDir, semester)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(academicDisciplines)
}

// GetAcademicDisciplineByEducationalDirectionAndName
// @Tags discipline
// @Summary      Get discipline by ed dir and name
// @Accept       json
// @Produce      json
// @Param ed_dir path string true "ed_dir"
// @Param name path string true "name"
// @Success 200 {object} entities.AcademicDiscipline
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/discipline/name/{ed_dir}/{name} [get]
// @Security ApiKeyAuth
func (h *Handler) GetAcademicDisciplineByEducationalDirectionAndName(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль админа
	edDir := c.Params("ed_dir")
	decodedEdDir, err := url.QueryUnescape(edDir)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid educational direction")
	}

	name := c.Params("name")
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid name")
	}

	h.logger.Debug().Msg("call h.services.AcademicDisciplineService.GetByEducationalDirectionAndName")
	fmt.Println(decodedName, decodedEdDir)
	academicDisciplines, err := h.services.AcademicDisciplineService.GetByEducationalDirectionAndName(c.Context(), decodedName, decodedEdDir)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(academicDisciplines)
}

// UpdateAcademicDiscipline
// @Tags discipline
// @Summary      Update discipline
// @Accept       json
// @Produce      json
// @Param data body entities.AcademicDiscipline true "university data"
// @Success 200 {object} entities.UpdateDeleteUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/discipline [put]
// @Security ApiKeyAuth
func (h *Handler) UpdateAcademicDiscipline(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль админа
	var academicDiscipline entities.AcademicDiscipline
	err := c.BodyParser(&academicDiscipline)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.AcademicDisciplineService.Update")
	err = h.services.AcademicDisciplineService.Update(c.Context(), &academicDiscipline)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("academicDiscipline with id=%v updated successfully", academicDiscipline.Id)})
}

// DeleteAcademicDiscipline
// @Tags discipline
// @Summary      Delete discipline
// @Accept       json
// @Produce      json
// @Param id path string true "discipline id"
// @Success 200 {object} entities.UpdateDeleteUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/discipline/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) DeleteAcademicDiscipline(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль админа
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid id in path")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id in path"})
	}

	h.logger.Debug().Msg("call h.services.AcademicDisciplineService.Delete")
	err = h.services.AcademicDisciplineService.Delete(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("academicDisciplineService with id=%v deleted successfully", id)})
}
