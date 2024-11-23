package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"os"
	"server/internal/entities"
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
// @Router       /auth/schedule/group/{group} [get]
// @Security ApiKeyAuth
func (h *Handler) GetScheduleByGroup(c *fiber.Ctx) error {
	groupName := c.Params("group")
	decodedGroupName, err := url.QueryUnescape(groupName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid group name")
	}

	h.logger.Debug().Msg("call h.services.ClassService.GetByGroupName")
	classes, err := h.services.ClassService.GetByGroupName(c.Context(), decodedGroupName)
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
// @Router       /auth/schedule/teacher/{teacher} [get]
// @Security ApiKeyAuth
func (h *Handler) GetScheduleByTeacher(c *fiber.Ctx) error {
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}
	h.logger.Debug().Msg("call h.services.UniversityService.GetByUserID")
	university, err := h.services.UniversityService.GetByUserID(c.Context(), userId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	teacherName := c.Params("teacher")
	decodedTeacherName, err := url.QueryUnescape(teacherName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid teacher name")
	}

	h.logger.Debug().Msg("call h.services.ClassService.GetByTeacherName")
	class, err := h.services.ClassService.GetByTeacherName(c.Context(), decodedTeacherName, university.Name)
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
// @Router     /auth/schedule/name/{name} [get]
// @Security ApiKeyAuth
func (h *Handler) GetScheduleByName(c *fiber.Ctx) error {
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}
	h.logger.Debug().Msg("call h.services.UniversityService.GetByUserID")
	university, err := h.services.UniversityService.GetByUserID(c.Context(), userId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	name := c.Params("name")
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid teacher name")
	}

	h.logger.Debug().Msg("call h.services.ClassService.GetByTeacherName")
	class, err := h.services.ClassService.GetByName(c.Context(), decodedName, university.Name)
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
// @Router       /auth/schedule/parse [post]
// @Security ApiKeyAuth
func (h *Handler) ParseSchedule(c *fiber.Ctx) error {
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}
	h.logger.Debug().Msg("call h.services.UniversityService.GetByUserID")
	university, err := h.services.UniversityService.GetByUserID(c.Context(), userId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	path := "./backend/tmp"
	err = os.MkdirAll(path, 0777)
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

		err = h.services.ClassService.Parse(randomFileName, university.Name)
		if err != nil {
			serviceError := err.Error()
			err = os.Remove(path + "/" + randomFileName)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": serviceError})
		}
		err = os.Remove(path + "/" + randomFileName)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

// GetScheduleSearchGroup
// @Tags schedule
// @Summary      Search groups
// @Accept       json
// @Produce      json
// @Success 200 {object} entities.ScheduleGroups
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/schedule/search/group [get]
// @Security ApiKeyAuth
func (h *Handler) GetScheduleSearchGroup(c *fiber.Ctx) error {
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}
	h.logger.Debug().Msg("call h.services.UniversityService.GetByUserID")
	university, err := h.services.UniversityService.GetByUserID(c.Context(), userId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call h.services.ClassService.SearchGroups")
	groups, err := h.services.ClassService.SearchGroups(c.Context(), university.Name)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(entities.ScheduleGroups{Groups: groups})
}

// GetScheduleSearchTeacher
// @Tags schedule
// @Summary      Search teachers
// @Accept       json
// @Produce      json
// @Success 200 {object} entities.ScheduleTeachers
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/schedule/search/teacher [get]
// @Security ApiKeyAuth
func (h *Handler) GetScheduleSearchTeacher(c *fiber.Ctx) error {
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}
	h.logger.Debug().Msg("call h.services.UniversityService.GetByUserID")
	university, err := h.services.UniversityService.GetByUserID(c.Context(), userId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call h.services.ClassService.SearchGroups")
	teachers, err := h.services.ClassService.SearchTeachers(c.Context(), university.Name)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(entities.ScheduleTeachers{Teachers: teachers})
}

// GetScheduleSearchName
// @Tags schedule
// @Summary      Search names
// @Accept       json
// @Produce      json
// @Success 200 {object} entities.ScheduleNames
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/schedule/search/name [get]
// @Security ApiKeyAuth
func (h *Handler) GetScheduleSearchName(c *fiber.Ctx) error {
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}
	h.logger.Debug().Msg("call h.services.UniversityService.GetByUserID")
	university, err := h.services.UniversityService.GetByUserID(c.Context(), userId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call h.services.ClassService.SearchGroups")
	names, err := h.services.ClassService.SearchNames(c.Context(), university.Name)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(entities.ScheduleNames{Names: names})
}
