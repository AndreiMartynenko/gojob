package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router //dependency
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	h.router.Get("/", h.home) //Method
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello, Fiber in web!")
}
