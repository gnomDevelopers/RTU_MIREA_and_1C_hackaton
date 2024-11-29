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

// GetScheduleOptionals
// @Tags schedule
// @Summary      Get schedule optionals
// @Accept       json
// @Produce      json
// @Param optionals path string true "schedule optionals"
// @Success 200 {object} []entities.Class
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router     /auth/schedule/optionals/{optionals} [get]
// @Security ApiKeyAuth
func (h *Handler) GetScheduleOptionals(c *fiber.Ctx) error {
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

	optionals := c.Params("optionals")
	decodedOptionals, err := url.QueryUnescape(optionals)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid teacher name")
	}

	h.logger.Debug().Msg("call h.services.ClassService.GetByTeacherName")
	class, err := h.services.ClassService.GetOptionals(c.Context(), decodedOptionals, university.Name)
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
	fmt.Printf("userId: %v", userId)
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
	if len(files) == 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg("empty file")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "empty file"})
	}
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

// ScheduleSearchGroup
// @Tags schedule
// @Summary      Search groups
// @Accept       json
// @Produce      json
// @Success 200 {object} []entities.ScheduleGroup
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/schedule/search/group [get]
// @Security ApiKeyAuth
func (h *Handler) ScheduleSearchGroup(c *fiber.Ctx) error {
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

	var scheduleGroups []entities.ScheduleGroup
	for _, group := range groups {
		scheduleGroups = append(scheduleGroups, entities.ScheduleGroup{Group: group})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(scheduleGroups)
}

// ScheduleSearchTeacher
// @Tags schedule
// @Summary      Search teachers
// @Accept       json
// @Produce      json
// @Success 200 {object} []entities.ScheduleTeachers
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/schedule/search/teacher [get]
// @Security ApiKeyAuth
func (h *Handler) ScheduleSearchTeacher(c *fiber.Ctx) error {
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

	var scheduleTeachers []entities.ScheduleTeachers
	for _, teacher := range teachers {
		scheduleTeachers = append(scheduleTeachers, entities.ScheduleTeachers{Teacher: teacher})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(scheduleTeachers)
}

// ScheduleSearchName
// @Tags schedule
// @Summary      Search names
// @Accept       json
// @Produce      json
// @Success 200 {object} []entities.ScheduleName
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/schedule/search/name [get]
// @Security ApiKeyAuth
func (h *Handler) ScheduleSearchName(c *fiber.Ctx) error {
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
	var scheduleNames []entities.ScheduleName
	for _, name := range names {
		scheduleNames = append(scheduleNames, entities.ScheduleName{Name: name})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(scheduleNames)
}

// GetGroupSubject
// @Tags schedule
// @Summary      Получение всех предметов группы
// @Accept       json
// @Produce      json
// @Success 200 {object} []entities.ScheduleName
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/schedule/group_subjects/{group} [get]
// @Security ApiKeyAuth
func (h *Handler) GetGroupSubject(c *fiber.Ctx) error {
	groupName := c.Params("group")
	decodedGroupName, err := url.QueryUnescape(groupName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid teacher name")
	}

	h.logger.Debug().Msg("call h.services.ClassService.SearchGroups")
	names, err := h.services.ClassService.SearchNamesWithGroup(c.Context(), decodedGroupName)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var scheduleNames []entities.ScheduleName
	for _, name := range names {
		scheduleNames = append(scheduleNames, entities.ScheduleName{Name: name})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(scheduleNames)
}
