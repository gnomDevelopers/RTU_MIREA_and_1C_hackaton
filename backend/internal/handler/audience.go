package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"server/internal/entities"
	"server/internal/log"
	"strconv"
)

// CreateAudiences
// @Tags audience
// @Summary      Create Audiences
// @Accept       json
// @Produce      json
// @Param data body []entities.CreateAudiencesRequest true "audience data"
// @Success 200 {object} []entities.CreateAudiencesResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /audience [post]
func (h *Handler) CreateAudiences(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	var audiences []entities.Audience
	err := c.BodyParser(&audiences)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.AudienceService.Create")
	ids, err := h.services.AudienceService.Create(c.Context(), &audiences)
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

// GetByIdAudience
// @Tags audience
// @Summary      Get audience by id
// @Accept       json
// @Produce      json
// @Param id path string true "audience id"
// @Success 200 {object} entities.Audience
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/audience/id/{id} [get]
func (h *Handler) GetByIdAudience(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	audience, err := h.services.AudienceService.GetById(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(audience)
}

// GetByCampusAudience
// @Tags audience
// @Summary      Get audience by name
// @Accept       json
// @Produce      json
// @Param name path string true "campus name"
// @Success 200 {object} entities.Audience
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/audience/campus/{name} [get]
func (h *Handler) GetByCampusAudience(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	name := c.Params("name")
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid campus name")
	}

	audience, err := h.services.AudienceService.GetByCampus(c.Context(), decodedName)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(audience)
}

// GetByTypeAudience
// @Tags audience
// @Summary      Get audience by type
// @Accept       json
// @Produce      json
// @Param type path string true "audience type"
// @Success 200 {object} entities.Audience
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/audience/type/{type} [get]
func (h *Handler) GetByTypeAudience(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	typeStr := c.Params("type")
	decodedType, err := url.QueryUnescape(typeStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid type")
	}

	audiences, err := h.services.AudienceService.GetByType(c.Context(), decodedType)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(audiences)
}

// GetByProfileAudience
// @Tags audience
// @Summary      Get audience by profile
// @Accept       json
// @Produce      json
// @Param profile path string true "audience profile"
// @Success 200 {object} entities.Audience
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/audience/profile/{profile} [get]
func (h *Handler) GetByProfileAudience(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	profile := c.Params("profile")
	decodedProfile, err := url.QueryUnescape(profile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid profile")
	}

	audiences, err := h.services.AudienceService.GetByProfile(c.Context(), decodedProfile)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(audiences)
}

// GetByCapacityAudience
// @Tags audience
// @Summary      Get audience by capacity
// @Accept       json
// @Produce      json
// @Param capacity path string true "audience capacity"
// @Success 200 {object} entities.Audience
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/audience/capacity/{capacity} [get]
func (h *Handler) GetByCapacityAudience(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	capacityStr := c.Params("capacity")
	capacity, err := strconv.Atoi(capacityStr)

	audiences, err := h.services.AudienceService.GetByCapacity(c.Context(), capacity)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(audiences)
}

// UpdateAudience
// @Tags audience
// @Summary      Update audience
// @Accept       json
// @Produce      json
// @Param data body entities.Audience true "audience data"
// @Success 200 {object} entities.UpdateDeleteAudienceResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/audience [put]
func (h *Handler) UpdateAudience(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	var audience entities.Audience
	err := c.BodyParser(&audience)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	err = h.services.AudienceService.Update(c.Context(), &audience)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("audience with id=%v updated successfully", audience.Id)})
}

// DeleteAudience
// @Tags audience
// @Summary      Delete audience
// @Accept       json
// @Produce      json
// @Param id path string true "audience id"
// @Success 200 {object} entities.UpdateDeleteAudienceResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router      /auth/audience/{id} [delete]
func (h *Handler) DeleteAudience(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid id in path")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id in path"})
	}

	err = h.services.AudienceService.Delete(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("audience with id=%v deleted successfully", id)})
}
