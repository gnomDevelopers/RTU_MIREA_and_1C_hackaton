package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
	"server/internal/entities"
	"server/internal/log"
	"server/pkg"
	"strconv"
	"strings"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var u entities.CreateUserRequest
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.services.UserService.CreateUser(c.Context(), &u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)

}

// Login
// @Tags         authentication
// @Summary      User login
// @Description  Authenticates a user and sets access and refresh tokens as cookies.
// @Accept       json
// @Produce      json
// @Param        data body entities.LoginUserRequest true "User login credentials"
// @Success      200 {object} entities.LoginUserResponse
// @Failure      400 {object} entities.ErrorResponse "Invalid request payload"
// @Failure      500 {object} entities.ErrorResponse "Internal server error"
// @Router       /login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var user entities.LoginUserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	u, err := h.services.UserService.Login(c.Context(), &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(u)
}

// GetUsersByUniversity
// @Tags user
// @Summary      Get users by university
// @Accept       json
// @Produce      json
// @Param university path string true "university"
// @Success 200 {object} []entities.User
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/user/university/{university} [get]
// @Security ApiKeyAuth
func (h *Handler) GetUsersByUniversity(c *fiber.Ctx) error {
	university := c.Params("university")
	decodedUniversity, err := url.QueryUnescape(university)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid university name")
	}

	h.logger.Debug().Msg("call h.services.UserService.GetByUniversity")
	class, err := h.services.UserService.GetByUniversity(c.Context(), decodedUniversity)
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

// CheckAuth
// @Tags         authentication
// @Summary      Authorization check
// @Description  Validates the JWT token from the Authorization header, extracts user ID, and generates a new access token.
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer JWT token"
// @Success      200 {object} map[string]interface{} "New access token and user ID"
// @Failure      400 {object} map[string]interface{} "Missing auth token"
// @Failure      401 {object} map[string]interface{} "Invalid auth header or token"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /login [get]
func (h *Handler) CheckAuth(c *fiber.Ctx) error {
	header := c.Get("Authorization")

	if header == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing auth token"})
	}

	tokenString := strings.Split(header, " ")
	if len(tokenString) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid auth header"})
	}

	token := tokenString[1]

	id, err := pkg.ParseToken(token, h.conf.Application.SigningKey)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	ExpirationTime, err := strconv.Atoi(h.conf.Application.TokenExpiration)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	newToken, err := pkg.GenerateAccessToken(id, ExpirationTime, h.conf.Application.SigningKey)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": newToken,
		"id":           id,
	})
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var user entities.CreateUserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req, err := h.services.UserService.CreateUser(c.Context(), &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}
	return c.Status(fiber.StatusOK).JSON(req)
}

// GetUserByID
// @Tags         user
// @Summary      Retrieve a user by ID
// @Description  Retrieves the user details based on the specified user ID provided in the URL parameter.
// @Accept       json
// @Produce      json
// @Param        id path integer true "User  ID"
// @Success      200 {object} entities.User "User  details"
// @Failure      400 {object} map[string]interface{} "Invalid user ID"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/user/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.services.UserService.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
