package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/url"
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
// @Router       /auth/campus/all [get]
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
// @Router       /auth/campus/id/{id} [get]
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
// @Router       /auth/campus/name/{name} [get]
func (h *Handler) GetByNameCampus(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	name := c.Params("name")
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid campus name")
	}

	campus, err := h.services.CampusService.GetByName(c.Context(), decodedName)
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
// @Router       /auth/campus/address/{address} [get]
func (h *Handler) GetByAddressCampus(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	address := c.Params("address")
	decodedAddress, err := url.QueryUnescape(address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid address")
	}
	fmt.Println(decodedAddress)
	campus, err := h.services.CampusService.GetByAddress(c.Context(), decodedAddress)
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

// GetByUniversityCampuses
// @Tags campus
// @Summary      Get campus by university
// @Accept       json
// @Produce      json
// @Param university path string true "campus university"
// @Success 200 {object} []entities.Campus
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/campus/university/{university} [get]
func (h *Handler) GetByUniversityCampuses(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	name := c.Params("university")
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid university name")
	}

	campuses, err := h.services.CampusService.GetByUniversity(c.Context(), decodedName)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(campuses)
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
// @Router       /auth/campus [put]
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
// @Router       /auth/campus/{id} [delete]
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
