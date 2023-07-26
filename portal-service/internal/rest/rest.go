package rest

import (
	"context"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/payment-system/portal-service/internal/service"
	"net/http"
)

var (
	svc service.Service
	app *fiber.App
)

func Start(service service.Service) error {
	svc = service

	app = fiber.New()

	public := app.Group("/")
	public.Get("/", welcome)

	private := app.Group("/v1")
	private.Get("/user", listUser)
	app.Add(http.MethodGet, "/", welcome)

	return app.Listen(":9001")
}

func Stop(ctx context.Context) {
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.WithContext(ctx).WithError(err).Fatal("Error shutting down rest server")
	}
}

func welcome(c *fiber.Ctx) error {
	return c.Status(200).JSON("WELCOME!!")
}

func listUser(c *fiber.Ctx) error {
	ctx := c.UserContext()
	user, err := svc.ListUsers(ctx)
	if err != nil {
		log.WithError(err).Error("error at retrieving users")
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(user)
}
