package handlers

import (
	"RMS_project/database"
	"RMS_project/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateJob(c *fiber.Ctx) error {
	var job models.Job

	if err := c.BodyParser(&job); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	job.PostedOn = time.Now()
	collection := database.DB.Collection("jobs")
	result, err := collection.InsertOne(context.Background(), job)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func GetJob(c *fiber.Ctx) error {
	jobID := c.Params("job_id")

	objectId, err := primitive.ObjectIDFromHex(jobID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid job ID"})
	}

	collection := database.DB.Collection("jobs")

	var job models.Job
	err = collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&job)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job not found"})
	}

	return c.JSON(job)
}

func GetJobs(c *fiber.Ctx) error {
	collection := database.DB.Collection("jobs")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer cursor.Close(context.Background())

	var jobs []models.Job
	if err := cursor.All(context.Background(), &jobs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(jobs)
}
