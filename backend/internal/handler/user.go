package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"server/internal/entities"
	"server/util"
	"strconv"
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

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    res.AccessToken,
		MaxAge:   60 * 60 * 6,
		Path:     "/",
		Domain:   ".gnomedeployer.ru",
		Secure:   false,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    res.RefreshToken,
		MaxAge:   60 * 60 * 724,
		Path:     "/",
		Domain:   ".gnomedeployer.ru",
		Secure:   false,
		HTTPOnly: true,
	})

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

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    u.AccessToken,
		MaxAge:   60 * 60 * 6,
		Path:     "/",
		Domain:   ".gnomedeployer.ru",
		Secure:   false,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    u.RefreshToken,
		MaxAge:   60 * 60 * 724,
		Path:     "/",
		Domain:   ".gnomedeployer.ru",
		Secure:   false,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(u)
}

// CheckAuth
// @Tags         authentication
// @Summary      Check user authentication
// @Description  Verifies the validity of the access token stored in cookies and retrieves the associated user ID.
// @Accept       json
// @Produce      json
// @Success      200 {object} entities.ErrorResponse "JWT is valid, user is authenticated"
// @Failure      401 {object} entities.ErrorResponse "Unauthorized, invalid or missing JWT"
// @Router       /login [get]
func (h *Handler) CheckAuth(c *fiber.Ctx) error {
	token := c.Cookies("access_token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"authorized": false,
			"message":    "JWT does not exist",
		})
	}

	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.conf.Application.SigningKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"authorized": false,
			"message":    "Invalid JWT",
		})
	}

	stringID, err := util.ExtractUserIDFromToken(token, h.conf.Application.SigningKey)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}
	userID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"authorized": true,
		"message":    "JWT is valid",
		"user_id":    userID,
	})
}

// Logout
// @Tags         authentication
// @Summary      User logout
// @Description  Logs out the user by clearing access and refresh tokens stored in cookies.
// @Accept       json
// @Produce      json
// @Success      200 {object} entities.LoginUserResponse "Logout successful message"
// @Failure      500 {object} entities.LoginUserResponse "Internal server error"
// @Router       /logout [post]
func (h *Handler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:   "access_token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
		Domain: ".gnomedeployer.ru",
		//Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:   "refresh_token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
		Domain: ".gnomedeployer.ru",
		//Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
	})

	// Возвращаем сообщение об успешном выходе
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "logout successful",
	})
}
