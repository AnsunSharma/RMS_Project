package handlers

import (
	"RMS_project/database"
	"RMS_project/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfile(c *fiber.Ctx) error {
	userID := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	collection := database.DB.Collection("profiles")

	var profile models.Profile
	err = collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&profile)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Profile not found"})
	}

	return c.JSON(profile)
}
