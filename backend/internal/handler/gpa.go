package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"server/internal/log"
	"strconv"
)

// GetByUserId
// @Tags gpa
// @Summary      Get gpa by user id
// @Accept       json
// @Produce      json
// @Param id path string true "user id"
// @Success 200 {object} entities.Class
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/gpa/id/{id} [get]
func (h *Handler) GetByUserId(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль проректора
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	h.logger.Debug().Msg("call h.services.GroupService.GetByUserID")
	group, err := h.services.GroupService.GetByUserID(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call h.services.UserData.GetEducationalDirection")
	edDir, err := h.services.UserService.GetEducationalDirection(c.Context(), id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call h.services.ClassService.SearchNamesWithGroup")
	classNames, err := h.services.ClassService.SearchNamesWithGroup(c.Context(), group.Name)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var disciplines []entities.AcademicDiscipline
	for _, className := range classNames {
		discipline, _ := h.services.AcademicDisciplineService.GetByEducationalDirectionAndName(c.Context(), className, edDir)

		if discipline != nil {
			disciplines = append(disciplines, *discipline)
		}
	}

	sumAllWorkHour := 0
	gpa := 0.0
	for _, discipline := range disciplines {
		score, _ := h.services.ScoreService.Get(c.Context(), &entities.Score{UserId: id, SubjectName: discipline.Name})
		if score != nil {
			disciplineHour := discipline.IndividualHours + discipline.PracticeHours + discipline.LectureHours + discipline.LabHours
			sumAllWorkHour += disciplineHour
			gpa += (float64(score.Sum) / float64(score.Count)) * float64(disciplineHour)
		}
	}
	gpa /= float64(sumAllWorkHour)
	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user_id": id, "gpa": gpa})
}
