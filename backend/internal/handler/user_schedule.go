package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"server/internal/log"
	"strconv"
)

// CreateUserSchedule
// @Tags user schedule
// @Summary      Create user schedule
// @Accept       json
// @Produce      json
// @Param data body entities.CreateUserScheduleRequest true "class data"
// @Success 200 {object} entities.CreateUserScheduleResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/user_schedule [post]
// @Security ApiKeyAuth
func (h *Handler) CreateUserSchedule(c *fiber.Ctx) error {
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}
	var userSchedule entities.UserSchedule
	err := c.BodyParser(&userSchedule)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	userSchedule.UserDataId = userId

	h.logger.Debug().Msg("call h.services.UserScheduleService.Create")
	id, err := h.services.UserScheduleService.Create(c.Context(), &userSchedule)
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

// GetUserSchedule
// @Tags user schedule
// @Summary      Get users schedule
// @Accept       json
// @Produce      json
// @Success 200 {object} entities.GetUserScheduleResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/user_schedule [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserSchedule(c *fiber.Ctx) error {
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}

	h.logger.Debug().Msg("call h.services.GroupService.GetByUserID")
	group, err := h.services.GroupService.GetByUserID(c.Context(), userId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call h.services.UserScheduleService.GetByUserId")
	userSchedule, err := h.services.UserScheduleService.GetByUserId(c.Context(), userId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call h.services.ClassService.GetByGroupName")
	groupSchedule, err := h.services.ClassService.GetByGroupName(c.Context(), group.Name)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(entities.GetUserScheduleResponse{UserSchedule: userSchedule, Classes: groupSchedule})
}

// UpdateUserSchedule
// @Tags user schedule
// @Summary      Update users schedule
// @Accept       json
// @Produce      json
// @Param data body entities.UpdateUserScheduleRequest true "users schedule data"
// @Success 200 {object} entities.UpdateDeleteUserScheduleResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/user_schedule [put]
// @Security ApiKeyAuth
func (h *Handler) UpdateUserSchedule(c *fiber.Ctx) error {
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}
	var userSchedule entities.UserSchedule
	err := c.BodyParser(&userSchedule)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	userSchedule.UserDataId = userId

	h.logger.Debug().Msg("call h.services.UserScheduleService.Update")
	err = h.services.UserScheduleService.Update(c.Context(), &userSchedule)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("users schedule with id=%v updated successfully", userSchedule.Id)})
}

// DeleteUserSchedule
// @Tags user schedule
// @Summary      Delete users schedule
// @Accept       json
// @Produce      json
// @Param id path string true "users schedule id"
// @Success 200 {object} entities.UpdateDeleteUserScheduleResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/user_schedule/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) DeleteUserSchedule(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid id in path")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id in path"})
	}

	h.logger.Debug().Msg("call h.services.UserScheduleService.Delete")
	err = h.services.UserScheduleService.Delete(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("users schedule with id=%v deleted successfully", id)})
}
