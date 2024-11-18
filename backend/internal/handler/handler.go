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

	f.Post("/sign-up", h.SignUp)
	f.Post("/login", h.Login)

	f.Get("/schedule", h.GetSchedule)

	f.Get("/university/all", h.GetAllUniversities)
	f.Get("/university/name/:name", h.GetByNameUniversity)
	f.Post("/university", h.CreateUniversity)
	f.Put("/university", h.UpdateUniversity)
	f.Delete("/university/:id", h.DeleteUniversity)

	f.Get("/campus/all", h.GetAllCampuses)
	f.Get("/campus/id/:id", h.GetByIdCampus)
	f.Get("/campus/name/:name", h.GetByNameCampus)
	f.Get("/campus/address/:address", h.GetByAddressCampus)
	f.Get("/campus/university_id/:id", h.GetByUniversityIdCampus)
	f.Post("/campus", h.CreateCampus)
	f.Put("/campus", h.UpdateCampus)
	f.Delete("/campus/:id", h.DeleteCampus)

	f.Get("/class/id/:id", h.GetByIdClass)
	f.Get("/class/group_name/:name", h.GetByGroupNameClass)
	f.Get("/class/teacher_name/:name", h.GetByTeacherNameClass)
	f.Get("/class/auditory_id/:id", h.GetByAuditoryIdClass)
	f.Post("/class", h.CreateClass)
	f.Put("/class", h.UpdateClass)
	f.Delete("/class/:id", h.DeleteClass)

	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return pkg.WithJWTAuth(c, h.conf.Application.SigningKey)
	})

	return f
}
