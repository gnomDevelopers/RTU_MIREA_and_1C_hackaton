package handler

import (
	_ "server/docs"
	"server/internal/config"
	"server/internal/log"
	"server/internal/service"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
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
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
	})

	f.Use(cors.New(cors.Config{
		//todo поменять на адрес фронта и раскомментить credentials (куки)
		AllowOrigins: "*",
		// AllowCredentials: true,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
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
	f.Get("/campus/university_id/:university_id", h.GetByUniversityIdCampus)
	f.Post("/campus", h.CreateCampus)
	f.Put("/campus", h.UpdateCampus)
	f.Delete("/campus/:id", h.DeleteCampus)

	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return pkg.WithJWTAuth(c, h.conf.Application.SigningKey)
	})

	return f
}
