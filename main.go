package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	fiberSwagger "github.com/arsmn/fiber-swagger/v2"

	_ "test/docs"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"test/route"

	database "test/configs"

	"test/app/repositories"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:2000
// @BasePath /
func main() {

	app := fiber.New()
	db := database.Init()
	repositories.DB = db
	app.Use(cors.New())
	app.Get("/swagger/*", fiberSwagger.Handler)

	route.SetApiRoutes(app)

	log.Fatal(app.Listen(":2000"))

}
