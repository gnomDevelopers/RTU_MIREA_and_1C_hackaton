package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
	"server/internal/entities"
)

// GetFacultiesByUniversityName
// @Tags         faculties
// @Summary      Retrieve faculties by university name
// @Description  Retrieves a list of faculties based on the specified university name provided in the URL parameter.
// @Accept       json
// @Produce      json
// @Param        group path string true "University name"
// @Success      200 {array} entities.Faculty "List of faculties for the specified university"
// @Failure      400 {object} map[string]interface{} "Invalid university name"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/faculties/university/{university} [get]
// @Security ApiKeyAuth
func (h *Handler) GetFacultiesByUniversityName(c *fiber.Ctx) error {
	universityName := c.Params("group")
	decodedUniversityName, err := url.QueryUnescape(universityName)
	req := entities.GetFacultyRequest{
		UniversityName: decodedUniversityName,
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res, err := h.services.FacultyService.GetAll(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

// CreateFaculty
// @Tags         faculties
// @Summary      Create a new faculty
// @Description  Creates a new faculty record in the system.
// @Accept       json
// @Produce      json
// @Param        body body entities.CreateFacultyRequest true "Faculty data"
// @Success      200 {object} entities.CreateFacultyResponse "Created faculty details"
// @Failure      400 {object} map[string]interface{} "Invalid request"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/faculties [post]
// @Security ApiKeyAuth
func (h *Handler) CreateFaculty(c *fiber.Ctx) error {
	var req entities.CreateFacultyRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.services.FacultyService.Create(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetAllFaculties
// @Tags         faculties
// @Summary      Retrieve all faculties
// @Description  Retrieves a list of all faculties based on optional filters provided in the request body.
// @Accept       json
// @Produce      json
// @Param        body body entities.GetFacultyRequest true "Request parameters for filtering faculties"
// @Success      200 {array} entities.Faculty "List of all faculties"
// @Failure      400 {object} map[string]interface{} "Invalid request data"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/faculties [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllFaculties(c *fiber.Ctx) error {
	var req entities.GetFacultyRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.services.FacultyService.GetAll(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
