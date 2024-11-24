package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog"
	_ "server/docs"
	"server/internal/config"
	"server/internal/log"
	"server/internal/service"
	"server/pkg"
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
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
	})

	f.Use(cors.New(cors.Config{
		AllowOrigins: "http://frontend:3000/",
		//AllowOrigins:     "http://localhost:8080/",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	f.Use(log.RequestLogger(h.logger))

	f.Get("/swagger/*", swagger.HandlerDefault)

	f.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("healthy")
	}) // healthcheck для докера

	f.Post("/sign-up", h.SignUp)
	f.Post("/login", h.Login)
	f.Get("/login", h.CheckAuth)
	f.Post("/logout", h.Logout)

	f.Post("/student/add", h.CreateStudent)

	f.Get("/university/all", h.GetAllUniversities)
	f.Get("/university/name/:name", h.GetByNameUniversity)
	f.Post("/university", h.CreateUniversity)
	f.Put("/university", h.UpdateUniversity)
	f.Delete("/university/:id", h.DeleteUniversity)

	f.Get("/campus/all", h.GetAllCampuses)
	f.Get("/campus/id/:id", h.GetByIdCampus)
	f.Get("/campus/name/:name", h.GetByNameCampus)
	f.Get("/campus/address/:address", h.GetByAddressCampus)
	f.Get("/campus/university/:university", h.GetByUniversityCampuses)
	f.Post("/campus", h.CreateCampuses)
	f.Put("/campus", h.UpdateCampus)
	f.Delete("/campus/:id", h.DeleteCampus)

	f.Get("/audience/id/:id", h.GetByIdAudience)
	f.Get("/audience/campus/:name", h.GetByCampusAudience)
	f.Get("/audience/type/:type", h.GetByTypeAudience)
	f.Get("/audience/profile/:profile", h.GetByProfileAudience)
	f.Get("/audience/capacity/:capacity", h.GetByCapacityAudience)
	f.Post("/audience", h.CreateAudiences)
	f.Put("/audience", h.UpdateAudience)
	f.Delete("/audience/:id", h.DeleteAudience)

	f.Get("/discipline/semester/:ed_dir/:semester", h.GetAcademicDisciplinesByEducationalDirectionAndSemester)
	f.Get("/discipline/name/:ed_dir/:name", h.GetAcademicDisciplineByEducationalDirectionAndName)
	f.Post("/discipline", h.CreateAcademicDiscipline)
	f.Put("/discipline", h.UpdateAcademicDiscipline)
	f.Delete("/discipline/:id", h.DeleteAcademicDiscipline)

	f.Get("/class/id/:id", h.GetByIdClass)
	f.Get("/class/auditory/:name", h.GetByAuditoryClass)
	f.Post("/class", h.CreateClasses)
	f.Put("/class", h.UpdateClass)
	f.Delete("/class/:id", h.DeleteClass)

	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return pkg.WithJWTAuth(c, h.conf.Application.SigningKey)
	})
	authGroup.Get("/user_schedule", h.GetUserSchedule)
	authGroup.Post("/user_schedule", h.CreateUserSchedule)
	authGroup.Put("/user_schedule", h.UpdateUserSchedule)
	authGroup.Delete("/user_schedule/:id", h.DeleteUserSchedule)

	authGroup.Post("/schedule/parse", h.ParseSchedule)
	authGroup.Get("/schedule/group/:group", h.GetScheduleByGroup)
	authGroup.Get("/schedule/teacher/:teacher", h.GetScheduleByTeacher)
	authGroup.Get("/schedule/name/:name", h.GetScheduleByName)

	authGroup.Get("/schedule/group_subjects", h.GetGroupSubject)
	authGroup.Post("/grade", h.CreateGrade)
	authGroup.Get("/grade/:group/:name", h.GetGradesBySubject)

	authGroup.Get("/schedule/search/teacher", h.GetScheduleSearchTeacher)
	authGroup.Get("/schedule/search/name", h.GetScheduleSearchName)
	authGroup.Get("/schedule/search/group", h.GetScheduleSearchGroup)

	return f
}
