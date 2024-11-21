package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"server/internal/log"
	"server/util"
)

// GetScheduleByGroup
// @Tags schedule
// @Summary      Get schedule
// @Accept       json
// @Produce      json
// @Param group path string true "schedule group"
// @Success 200 {object} []entities.Class
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /schedule/group/{group} [get]
func (h *Handler) GetScheduleByGroup(c *fiber.Ctx) error {
	groupName := c.Params("name")

	h.logger.Debug().Msg("call h.services.ClassService.GetByGroupName")
	classes, err := h.services.ClassService.GetByGroupName(c.Context(), groupName)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(classes)
}

// GetScheduleByTeacher
// @Tags schedule
// @Summary      Get schedule by teacher
// @Accept       json
// @Produce      json
// @Param teacher path string true "schedule teacher"
// @Success 200 {object} []entities.Class
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /schedule/teacher/{teacher} [get]
func (h *Handler) GetScheduleByTeacher(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	teacherName := c.Params("name")

	h.logger.Debug().Msg("call h.services.ClassService.GetByTeacherName")
	class, err := h.services.ClassService.GetByTeacherName(c.Context(), teacherName)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(class)
}

// GetScheduleByName
// @Tags schedule
// @Summary      Get schedule by name
// @Accept       json
// @Produce      json
// @Param name path string true "schedule name"
// @Success 200 {object} []entities.Class
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /schedule/name/{name} [get]
func (h *Handler) GetScheduleByName(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON("")
}

// ParseSchedule
// @Tags schedule
// @Summary      Get schedule by name
// @Accept       mpfd
// @Produce      json
// @Param file formData file true "Upload schedule file"
// @Success 200 {object} []entities.ParseScheduleResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /schedule/parse [post]
func (h *Handler) ParseSchedule(c *fiber.Ctx) error {
	path := "./backend/tmp"
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	files := form.File["file"]
	for _, file := range files {
		if file.Header["Content-Type"][0] != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "unsupported file extension"})
		}
		randomFileName := util.GenerateRandomString(20) + ".xlsx"
		file.Filename = randomFileName

		if err = c.SaveFile(file, fmt.Sprintf("./backend/tmp/%s", file.Filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		err = h.services.ClassService.Parse(randomFileName)
		if err != nil {
			serviceError := err.Error()
			err = os.Remove(path + "/" + randomFileName)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": serviceError})
		}
		err = os.Remove(path + "/" + randomFileName)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}
