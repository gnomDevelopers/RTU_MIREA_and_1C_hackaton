package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"server/internal/entities"
	"server/internal/log"
)

// CreateGrade
// @Tags grade
// @Summary      Create grade
// @Accept       json
// @Produce      json
// @Param data body entities.CreateGradeRequest true "grade data"
// @Success 200 {object} entities.CreateGradeResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/grade [post]
// @Security ApiKeyAuth
func (h *Handler) CreateGrade(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль
	var grade entities.Grade
	err := c.BodyParser(&grade)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call h.services.CampusService.Create")
	id, err := h.services.GradeService.Create(c.Context(), &grade)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	class, err := h.services.ClassService.GetById(c.Context(), grade.ClassId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.services.ScoreService.Update(c.Context(), &entities.Score{UserId: grade.UserId, Sum: grade.Value, Count: 1, SubjectName: class.Name})
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(id)
}

// GetGradesBySubject
// @Tags grade
// @Summary      Get grade by group
// @Accept       json
// @Produce      json
// @Param group path string true "group"
// @Param name path string true "name"
// @Success 200 {object} []entities.GetGradesBySubject
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/grade/{group}/{name} [get]
// @Security ApiKeyAuth
func (h *Handler) GetGradesBySubject(c *fiber.Ctx) error {
	// TODO: добавить проверку на роль
	group := c.Params("group")
	decodedGroup, err := url.QueryUnescape(group)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid group")
	}
	name := c.Params("name")
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid name")
	}

	groupMember, err := h.services.GroupService.GetGroupMembers(c.Context(), decodedGroup)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	classes, err := h.services.ClassService.GetByNameAndGroupWithoutLk(c.Context(), decodedName, decodedGroup)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Println("OK")
	for i := range *classes {
		if (*classes)[i].Grades == nil {
			(*classes)[i].Grades = []entities.Grade{}
		}

		for _, member := range *groupMember {
			grade, _ := h.services.GradeService.GetByUserIdAndClassId(c.Context(), member.ID, (*classes)[i].Id)
			if grade != nil {
				(*classes)[i].Grades = append((*classes)[i].Grades, *grade)
			}
		}
	}

	var getGradesBySubject entities.GetGradesBySubject
	getGradesBySubject.GradeClass = *classes
	getGradesBySubject.GroupMember = *groupMember

	for _, member := range *groupMember {
		var averageUsersScore entities.UsersScore
		score, _ := h.services.ScoreService.Get(c.Context(), &entities.Score{UserId: member.ID, SubjectName: decodedName})
		if score == nil {
			continue
		}
		averageUsersScore.UserId = member.ID
		averageUsersScore.SumScore = score.Sum
		averageUsersScore.AverageScore = float64(score.Sum) / float64(score.Count)
		getGradesBySubject.UsersScore = append(getGradesBySubject.UsersScore, averageUsersScore)
	}

	return c.Status(fiber.StatusOK).JSON(getGradesBySubject)
}
