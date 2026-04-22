package home

import (
	"gojob/pkg/tadapter"
	"gojob/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router //dependency
	customLogger *zerolog.Logger
}

// api routers group
// /v1/api/
// ?v1/api/error

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}

	// api := h.router.Group("/api")
	h.router.Get("/", h.home)
	h.router.Get("/error", h.error)

	//h.router.Get("/", h.home) //Method
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	//component.Render(context.Background(), os.Stdout)
	//return c.SendString("sd")
	return tadapter.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {

	//levels of logs
	//log.Trace("Trace")
	//log.Debug("Debug")
	//log.Info("Info")
	//log.Warn("Warn")
	//log.Error("Error")
	//log.Fatal("Fatal")
	//log.Panic("Panic")
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.com").
		Int("Id", 23).
		Msg("Info")

	//Custom logger for specific tasks
	//logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(1)
	//logger.Info().Msg("Logger 2")

	return c.SendString("Error!")
}
