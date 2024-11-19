package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"server/internal/log"
	"strconv"
)

// CreateClasses
// @Tags class
// @Summary      Create classes
// @Accept       json
// @Produce      json
// @Param data body []entities.CreateClassesRequest true "class data"
// @Success 200 {object} []entities.CreateClassesResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /class [post]
func (h *Handler) CreateClasses(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	var classes []entities.Class
	err := c.BodyParser(&classes)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.ClassService.Create")
	ids, err := h.services.ClassService.Create(c.Context(), &classes)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(ids)
}

// GetByIdClass
// @Tags class
// @Summary      Get class by id
// @Accept       json
// @Produce      json
// @Param id path string true "class id"
// @Success 200 {object} entities.Class
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /class/id/{id} [get]
func (h *Handler) GetByIdClass(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	h.logger.Debug().Msg("call h.services.ClassService.GetById")
	class, err := h.services.ClassService.GetById(c.Context(), id)
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

// GetByGroupNameClass
// @Tags class
// @Summary      Get class by group name
// @Accept       json
// @Produce      json
// @Param id path string true "class group name"
// @Success 200 {object} entities.Class
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /class/group_name/{name} [get]
func (h *Handler) GetByGroupNameClass(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	groupName := c.Params("name")

	h.logger.Debug().Msg("call h.services.ClassService.GetByGroupName")
	class, err := h.services.ClassService.GetByGroupName(c.Context(), groupName)
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

// GetByTeacherNameClass
// @Tags class
// @Summary      Get class by teacher name
// @Accept       json
// @Produce      json
// @Param id path string true "class teacher name"
// @Success 200 {object} entities.Class
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /class/teacher_name/{name} [get]
func (h *Handler) GetByTeacherNameClass(c *fiber.Ctx) error {
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

// GetByAuditoryIdClass
// @Tags class
// @Summary      Get class by auditory id
// @Accept       json
// @Produce      json
// @Param id path string true "class auditory id"
// @Success 200 {object} entities.Class
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /class/auditory_id/{id} [get]
func (h *Handler) GetByAuditoryIdClass(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	auditoryIdStr := c.Params(":id")
	auditoryId, err := strconv.Atoi(auditoryIdStr)

	h.logger.Debug().Msg("call h.services.ClassService.GetByAuditoryId")
	class, err := h.services.ClassService.GetByAuditoryId(c.Context(), auditoryId)
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

// UpdateClass
// @Tags class
// @Summary      Update class
// @Accept       json
// @Produce      json
// @Param data body entities.Class true "class data"
// @Success 200 {object} entities.UpdateDeleteClassResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /class [put]
func (h *Handler) UpdateClass(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	var class entities.Class
	err := c.BodyParser(&class)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.ClassService.Update")
	err = h.services.ClassService.Update(c.Context(), &class)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("class with id=%v updated successfully", class.Id)})
}

// DeleteClass
// @Tags class
// @Summary      Delete class
// @Accept       json
// @Produce      json
// @Param id path string true "class id"
// @Success 200 {object} entities.UpdateDeleteClassResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /class{id} [delete]
func (h *Handler) DeleteClass(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid id in path")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id in path"})
	}

	h.logger.Debug().Msg("call h.services.ClassService.Delete")
	err = h.services.ClassService.Delete(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("class with id=%v deleted successfully", id)})
}
