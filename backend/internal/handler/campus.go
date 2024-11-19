package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"server/internal/log"
	"strconv"
)

// CreateCampuses
// @Tags campus
// @Summary      Create campuses
// @Accept       json
// @Produce      json
// @Param data body []entities.CreateCampusesRequest true "campus data"
// @Success 200 {object} []entities.CreateCampusesResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /campus [post]
func (h *Handler) CreateCampuses(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	var campuses []entities.Campus
	err := c.BodyParser(&campuses)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.CampusService.Create")
	ids, err := h.services.CampusService.Create(c.Context(), &campuses)
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

// GetAllCampuses
// @Tags campus
// @Summary      Get all campuses
// @Accept       json
// @Produce      json
// @Success 200 {object} []entities.Campus
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /campus/all [get]
func (h *Handler) GetAllCampuses(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	campuses, err := h.services.CampusService.GetAll(c.Context())
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(campuses)
}

// GetByIdCampus
// @Tags campus
// @Summary      Get campus by id
// @Accept       json
// @Produce      json
// @Param id path string true "campus id"
// @Success 200 {object} entities.Campus
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /campus/id/{id} [get]
func (h *Handler) GetByIdCampus(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	campus, err := h.services.CampusService.GetById(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(campus)
}

// GetByNameCampus
// @Tags campus
// @Summary      Get campus by name
// @Accept       json
// @Produce      json
// @Param name path string true "campus name"
// @Success 200 {object} entities.Campus
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /campus/name/{name} [get]
func (h *Handler) GetByNameCampus(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	name := c.Params("name")

	campus, err := h.services.CampusService.GetByName(c.Context(), name)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(campus)
}

// GetByAddressCampus
// @Tags campus
// @Summary      Get campus by address
// @Accept       json
// @Produce      json
// @Param address path string true "campus address"
// @Success 200 {object} entities.Campus
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /campus/address/{address} [get]
func (h *Handler) GetByAddressCampus(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	address := c.Params("address")

	campus, err := h.services.CampusService.GetByAddress(c.Context(), address)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(campus)
}

// GetByUniversityIdCampus
// @Tags campus
// @Summary      Get campus by university_id
// @Accept       json
// @Produce      json
// @Param university_id path string true "campus university_id"
// @Success 200 {object} entities.Campus
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /campus/university_id/{id} [get]
func (h *Handler) GetByUniversityIdCampus(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	universityIdStr := c.Params("id")
	universityId, err := strconv.Atoi(universityIdStr)

	campus, err := h.services.CampusService.GetByUniversityId(c.Context(), universityId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(campus)
}

// UpdateCampus
// @Tags campus
// @Summary      Update campus
// @Accept       json
// @Produce      json
// @Param data body entities.Campus true "campus data"
// @Success 200 {object} entities.UpdateDeleteCampusResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /campus [put]
func (h *Handler) UpdateCampus(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	var campus entities.Campus
	err := c.BodyParser(&campus)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	err = h.services.CampusService.Update(c.Context(), &campus)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("campus with id=%v updated successfully", campus.Id)})
}

// DeleteCampus
// @Tags campus
// @Summary      Delete campus
// @Accept       json
// @Produce      json
// @Param id path string true "campus id"
// @Success 200 {object} entities.UpdateDeleteCampusResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /campus{id} [delete]
func (h *Handler) DeleteCampus(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid id in path")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id in path"})
	}

	err = h.services.CampusService.Delete(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("campus with id=%v deleted successfully", id)})
}
