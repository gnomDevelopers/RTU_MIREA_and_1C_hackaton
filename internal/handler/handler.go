package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
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
		//todo поменять на адрес фронта и раскомментить credentials (куки)
		AllowOrigins: "*",
		// AllowCredentials: true,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	f.Use(log.RequestLogger(h.logger))

	f.Post("/sign-up", h.SignUp)
	f.Post("/login", h.Login)

	f.Get("/schedule", h.GetSchedule)

	f.Get("/university/all", h.GetAllUniversities)
	f.Get("/university/:name", h.GetByNameUniversity)
	f.Post("/university", h.CreateUniversity)
	f.Put("/university", h.UpdateUniversity)
	f.Delete("/university", h.DeleteUniversity)

	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return pkg.WithJWTAuth(c, h.conf.Application.SigningKey)
	})

	return f
}
