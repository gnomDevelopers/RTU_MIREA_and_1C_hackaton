package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"server/internal/log"
	"strconv"
)

// ExistsWorkUser
// @Tags work
// @Summary      Login in hr
// @Accept       json
// @Produce      json
// @Param id path string true "university id"
// @Success 200 {object} entities.ExistsResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /work/exists/id/{id} [post]
// @Security ApiKeyAuth
func (h *Handler) ExistsWorkUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	h.logger.Debug().Msg("call h.services.WorkService.Exists")
	flag, err := h.services.WorkService.Exists(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"exists": flag})
}

// LoginHR
// @Tags         work
// @Summary      Hr login
// @Description  Authenticates a user and sets access and refresh tokens as cookies.
// @Accept       json
// @Produce      json
// @Param        data body entities.LoginUserRequest true "User login credentials"
// @Success      200 {object} entities.LoginUserResponse
// @Failure      400 {object} entities.ErrorResponse "Invalid request payload"
// @Failure      500 {object} entities.ErrorResponse "Internal server error"
// @Router       /work/login/hr [post]
func (h *Handler) LoginHR(c *fiber.Ctx) error {
	var user entities.LoginUserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	h.logger.Debug().Msg("call h.services.WorkService.Login")
	u, err := h.services.WorkService.Login(c.Context(), &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(u)
}

// GetByIdWorkUserProfile
// @Tags work
// @Summary      Update profile
// @Accept       json
// @Produce      json
// @Param id path string true "university id"
// @Success 200 {object} entities.WorkUser
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/work/profile/id/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) GetByIdWorkUserProfile(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	h.logger.Debug().Msg("call h.services.WorkService.UpdateWorkUser")
	workUser, err := h.services.WorkService.GetByIdWorkUser(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(workUser)
}

// UpdateWorkUserProfile
// @Tags work
// @Summary      Update profile
// @Accept       json
// @Produce      json
// @Param data body entities.WorkUserUpdateRequest true "profile data"
// @Success 200 {object} entities.WorkUserUpdateResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/work/profile [put]
// @Security ApiKeyAuth
func (h *Handler) UpdateWorkUserProfile(c *fiber.Ctx) error {
	var workUser entities.WorkUserUpdateRequest
	err := c.BodyParser(&workUser)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.WorkService.UpdateWorkUser")
	err = h.services.WorkService.UpdateWorkUser(c.Context(), &workUser)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("university with id=%v updated successfully", workUser.Id)})
}

// ResponseCandidate
// @Tags work
// @Summary      Response
// @Accept       json
// @Produce      json
// @Param data body entities.CreateResponse true "response data"
// @Success 200 {object} entities.WorkUserUpdateResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/work/response [post]
// @Security ApiKeyAuth
func (h *Handler) ResponseCandidate(c *fiber.Ctx) error {
	var response entities.Response
	err := c.BodyParser(&response)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.WorkService.CreateResponse")
	err = h.services.WorkService.CreateResponse(c.Context(), &response)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

// GetCandidateResponses
// @Tags work
// @Summary      Response
// @Accept       json
// @Produce      json
// @Param id path string true "candidate id"
// @Success 200 {object} entities.CreateUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/work/response/candidate/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) GetCandidateResponses(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	h.logger.Debug().Msg("call h.services.WorkService.GetWorkUserResponses")
	university, err := h.services.WorkService.GetWorkUserResponses(c.Context(), id)
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

// GetHRResponses
// @Tags work
// @Summary      Response
// @Accept       json
// @Produce      json
// @Param id path string true "candidate id"
// @Success 200 {object} entities.CreateUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/work/response/hr/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) GetHRResponses(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	h.logger.Debug().Msg("call h.services.WorkService.GetHRResponses")
	university, err := h.services.WorkService.GetHRResponses(c.Context(), id)
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

// GetAllWorkUserId
// @Tags work
// @Summary      Response
// @Accept       json
// @Produce      json
// @Param id path string true "candidate id"
// @Success 200 {object} entities.CreateUniversityResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/work/response/hr/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllWorkUserId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	h.logger.Debug().Msg("call h.services.WorkService.GetHRResponses")
	university, err := h.services.WorkService.GetHRResponses(c.Context(), id)
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
