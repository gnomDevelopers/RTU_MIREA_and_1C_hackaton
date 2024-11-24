package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
)

// CreateStudent
// @Tags         user data
// @Summary      Create student
// @Description  Creates multiple students with their associated data (university, faculty, department, group).
// @Accept       json
// @Produce      json
// @Param        data body []entities.AddUserDataRequest true "Array of user data to be added"
// @Success      200 {array} []entities.AddUserDataResponse
// @Failure      400 {object} entities.ErrorResponse "Invalid request payload"
// @Failure      500 {object} entities.ErrorResponse "Internal server error"
// @Router       /student/add [post]
func (h *Handler) CreateStudent(c *fiber.Ctx) error {
	var requests []entities.AddUserDataRequest

	// Парсим JSON из запроса
	if err := c.BodyParser(&requests); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Вызываем метод сервиса
	resp, err := h.services.UserData.AddStudent(c.Context(), &requests)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Возвращаем успешный ответ
	return c.Status(fiber.StatusOK).JSON(resp)
}
