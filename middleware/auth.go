package middlewares

import (
	"RMS_project/database"
	"RMS_project/models"
	"RMS_project/utils"
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get("Authorization")
		if authorizationHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Split the "Bearer" prefix from the token
		tokenParts := strings.Split(authorizationHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		token := tokenParts[1]
		_, err := utils.ParseToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		return c.Next()
	}
}

func IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Split the "Bearer" prefix from the token
		roleId := c.Query("role_id")
		objID, err := primitive.ObjectIDFromHex(roleId)
		if err != nil {
			fmt.Println("Error converting string to ObjectId:", err)
		}
		collection := database.DB.Collection("users")
		var user models.User
		err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
		if err != nil {
			fmt.Println("Error finding user:", err)
		}
		fmt.Println(user)
		if user.UserType != "Admin" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user is not allow to get data"})
		}
		return c.Next()
	}
}
