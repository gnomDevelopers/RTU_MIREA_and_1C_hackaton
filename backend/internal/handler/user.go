package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
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
		MaxAge:   60 * 60 * 1,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    res.RefreshToken,
		MaxAge:   60 * 60 * 24 * 90,
		Path:     "/",
		Domain:   "localhost",
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
		MaxAge:   60 * 60 * 1,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    u.RefreshToken,
		MaxAge:   60 * 60 * 24 * 90,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(u)
}
