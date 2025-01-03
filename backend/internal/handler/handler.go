package handler

import (
	_ "server/docs"
	"server/internal/config"
	"server/internal/log"
	"server/internal/service"
	"server/pkg"

	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
)

type Handler struct {
	services *service.Service
	logger   *zerolog.Logger
	conf     *config.Config
}

func NewHandler(services *service.Service, logger *zerolog.Logger, conf *config.Config) *Handler {
	return &Handler{services: services, logger: logger, conf: conf}
}

func (h *Handler) Router() *fiber.App {
	f := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})

	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		//AllowOrigins:     "http://localhost:8080/",
		//AllowCredentials: true,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	f.Use(log.RequestLogger(h.logger))

	f.Get("/swagger/*", fiberSwagger.WrapHandler)

	f.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("healthy")
	}) // healthcheck для докера

	f.Post("/login", h.Login)
	f.Get("/login", h.CheckAuth)

	f.Post("/work/login/hr", h.LoginHR)
	f.Get("/work/exists/id/:id", h.ExistsWorkUser)
	//f.Get("/user/:id", h.GetUserByID)

	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return pkg.WithJWTAuth(c, h.conf.Application.SigningKey)
	})

	authGroup.Get("/user/university/:university", h.GetUsersByUniversity)
	authGroup.Get("/user/:id", h.GetUserByID)
	authGroup.Post("/user", h.CreateUsers)
	authGroup.Put("/user/:id/promote", h.UpdateUserRole)

	authGroup.Get("/user_schedule", h.GetUserSchedule)
	authGroup.Post("/user_schedule", h.CreateUserSchedule)
	authGroup.Put("/user_schedule", h.UpdateUserSchedule)
	authGroup.Delete("/user_schedule/:id", h.DeleteUserSchedule)

	authGroup.Post("/schedule/parse", h.ParseSchedule)
	authGroup.Get("/schedule/group/:group", h.GetScheduleByGroup)
	authGroup.Get("/schedule/teacher/:teacher", h.GetScheduleByTeacher)
	authGroup.Get("/schedule/optionals/:optionals", h.GetScheduleOptionals)
	authGroup.Get("/schedule/group_subjects/:group", h.GetGroupSubject)

	authGroup.Delete("/achievement/:id", h.DeleteAchievement)
	authGroup.Post("/achievement", h.CreateAchievement)
	authGroup.Get("/achievement/user/:id", h.GetAchievementByUserID)

	authGroup.Post("/grade", h.CreateGrade)
	authGroup.Get("/grade/:group/:name", h.GetGradesBySubject)

	authGroup.Get("/schedule/search/teacher", h.ScheduleSearchTeacher)
	authGroup.Get("/schedule/search/name", h.ScheduleSearchName)
	authGroup.Get("/schedule/search/group", h.ScheduleSearchGroup)

	authGroup.Post("/work/response", h.ResponseCandidate)
	authGroup.Get("/work/response/candidate/:id", h.GetCandidateResponses)

	// Перенесено
	authGroup.Post("/group", h.CreateGroup)
	authGroup.Get("/groups/university/:university", h.GetAllGroups)

	authGroup.Get("/university/all", h.GetAllUniversities)
	authGroup.Get("/university/name/:name", h.GetByNameUniversity)
	authGroup.Get("/university/id/:id", h.GetByIdUniversity)
	authGroup.Post("/university", h.CreateUniversity)
	authGroup.Put("/university", h.UpdateUniversity)
	authGroup.Delete("/university/:id", h.DeleteUniversity)

	authGroup.Post("/faculty", h.CreateFaculty)
	authGroup.Get("/faculties/university/:university", h.GetFacultiesByUniversityName)

	authGroup.Get("/campus/all", h.GetAllCampuses)
	authGroup.Get("/campus/id/:id", h.GetByIdCampus)
	authGroup.Get("/campus/name/:name", h.GetByNameCampus)
	authGroup.Get("/campus/address/:address", h.GetByAddressCampus)
	authGroup.Get("/campus/university/:university", h.GetByUniversityCampuses)
	authGroup.Post("/campus", h.CreateCampuses)
	authGroup.Put("/campus", h.UpdateCampus)
	authGroup.Delete("/campus/:id", h.DeleteCampus)

	authGroup.Get("/audience/id/:id", h.GetByIdAudience)
	authGroup.Get("/audience/campus/:name", h.GetByCampusAudience)
	authGroup.Get("/audience/type/:type", h.GetByTypeAudience)
	authGroup.Get("/audience/profile/:profile", h.GetByProfileAudience)
	authGroup.Get("/audience/capacity/:capacity", h.GetByCapacityAudience)
	authGroup.Get("/audience/university/:university", h.GetByUniversityAudiences)
	authGroup.Post("/audience", h.CreateAudiences)
	authGroup.Put("/audience", h.UpdateAudience)
	authGroup.Delete("/audience/:id", h.DeleteAudience)

	authGroup.Get("/discipline/semester/:ed_dir/:semester", h.GetAcademicDisciplinesByEducationalDirectionAndSemester)
	authGroup.Get("/discipline/name/:ed_dir/:name", h.GetAcademicDisciplineByEducationalDirectionAndName)
	authGroup.Post("/discipline", h.CreateAcademicDiscipline)
	authGroup.Put("/discipline", h.UpdateAcademicDiscipline)
	authGroup.Delete("/discipline/:id", h.DeleteAcademicDiscipline)

	authGroup.Get("/class/id/:id", h.GetByIdClass)
	authGroup.Get("/class/auditory/:name", h.GetByAuditoryClass)
	authGroup.Post("/class", h.CreateClasses)
	authGroup.Put("/class", h.UpdateClass)
	authGroup.Delete("/class/:id", h.DeleteClass)
	authGroup.Get("/class/:id/participants", h.GetAllClassParticipants)
	authGroup.Get("/class/:id/visiting", h.GetAllClassVisiting)
	authGroup.Post("/class/visiting", h.AddClassVisiting)
	authGroup.Post("/class/visiting/check-in", h.CheckIn)

	authGroup.Get("/gpa/id/:id", h.GetGpaByUserId)
	authGroup.Put("/gpa", h.UpdateGpa)

	authGroup.Get("/percentile/id/:id", h.GetPercentileByUserId)

	authGroup.Get("/department/university/:university", h.GetByUniversityDepartments)
	authGroup.Post("/department", h.CreateDepartment)

	authGroup.Get("/work/exists/id/:id", h.ExistsWorkUser)
	authGroup.Get("/work/profile/id/:id", h.GetByIdWorkUserProfile)
	authGroup.Get("/work/response/candidate/id/:id", h.GetCandidateResponses)
	authGroup.Get("/work/response/hr/:id", h.GetHRResponses)
	authGroup.Get("/work/work_user/all", h.GetAllWorkUser)
	authGroup.Post("/work/login/hr", h.LoginHR)
	authGroup.Post("/work/response", h.ResponseCandidate)
	authGroup.Put("/work/profile", h.UpdateWorkUserProfile)

	return f
}
