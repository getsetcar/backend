package main

import (
	"getsetcar/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	var err error
	carsData, err := ReadCompressedJSON("./cars.json.gz")
	if err != nil {
		log.Fatalf("Failed to read compressed JSON: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "GetSetCar API",
	})

	// Add middleware
	app.Use(cors.New())

	// Initialize handlers
	carHandler := handlers.NewCarHandler(*carsData)

	// Register routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to GetSetCar API",
		})
	})
	app.Get("/cars/:brand", carHandler.GetCarsForBrand)
	app.Get("/cars/:brand/:model", carHandler.GetModel)

	// Start server
	log.Printf("Server starting on port 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
