package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jpmoraess/appointment-api/internal/handlers"
	"github.com/jpmoraess/appointment-api/pkg/metrics"
)

type Router struct {
	app             *fiber.App
	registerHandler *handlers.RegisterHandler
}

func NewRouter(app *fiber.App, registerHandler *handlers.RegisterHandler) *Router {
	return &Router{
		app:             app,
		registerHandler: registerHandler,
	}
}

func (r *Router) SetupRoutes() {
	auth := r.app.Group("/auth")
	auth.Post("/register", r.registerHandler.HandleRegister)

	// swagger
	r.app.Get("/swagger/*", swagger.HandlerDefault)
	r.app.Get("/swagger", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})

	// metrics
	r.app.Get("/metrics", metrics.PrometheusHandler)
}
