package main

import (
	"RMS_project/database"
	"RMS_project/handlers"
	middlewares "RMS_project/middleware"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	app := fiber.New()

	app.Use(logger.New())

	app.Post("/signup", handlers.SignUp)
	app.Post("/login", handlers.Login)

	app.Post("/uploadResume", middlewares.Protected(), handlers.UploadResume)

	app.Post("/admin/job", middlewares.Protected(), middlewares.IsAdmin(), handlers.CreateJob)

	app.Get("/admin/job/:job_id", middlewares.Protected(), middlewares.IsAdmin(), handlers.GetJob)
	// app.Get("/admin/applicants", middlewares.Protected(), handlers.GetApplicants)
	// app.Get("/admin/applicant/:applicant_id", middlewares.Protected(), handlers.GetApplicant)

	app.Get("/jobs", middlewares.Protected(), handlers.GetJobs)
	// app.Get("/jobs/apply", middlewares.Protected(), handlers.ApplyJob)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
