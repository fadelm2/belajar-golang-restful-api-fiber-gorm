package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"restful-api-gorm-fiber/config"
	"restful-api-gorm-fiber/controller"
	"restful-api-gorm-fiber/exception"
	"restful-api-gorm-fiber/model"
	"restful-api-gorm-fiber/repository"
	"restful-api-gorm-fiber/router"
	"restful-api-gorm-fiber/service"
)

func main() {
	fmt.Println("Run Service")

	db := config.OpenConnection()
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	//init Repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	//init Service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	//note contoller
	noteContoller := controller.NewNoteController(noteService)

	//routes
	routes := router.NewRouter(noteContoller)

	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	app.Use(logger.New())
	app.Mount("/api", routes)
	log.Fatal(app.Listen(":8000"))
}
