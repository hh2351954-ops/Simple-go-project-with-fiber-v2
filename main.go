package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Get("/about", func(c *fiber.Ctx) error {
        return c.SendString("About page")
    })

    app.Get("/contact", func(c *fiber.Ctx) error {
        return c.SendString("Contact page")
    })

    app.Post("/submit", func(c *fiber.Ctx) error {
        // Example: Accept JSON data
        var data map[string]interface{}
        if err := c.BodyParser(&data); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Invalid input",
            })
        }
        return c.JSON(data)
    })

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "OK"})
    })

    app.Listen(":3000")
}
