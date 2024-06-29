package handlers

import (
	"RMS_project/database"
	"RMS_project/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const apiEndpoint = "https://api.apilayer.com/resume_parser/upload"
const apiKey = "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1"

func UploadResume(c *fiber.Ctx) error {
	file, err := c.FormFile("resume")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to get file"})
	}

	if file.Header.Get("Content-Type") != "application/pdf" && file.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Only PDF and DOCX formats are allowed"})
	}

	userID := c.Locals("user_id").(string)

	fileContent, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
	}
	defer fileContent.Close()

	body, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read file"})
	}

	request, err := http.NewRequest("POST", apiEndpoint, bytes.NewReader(body))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create request"})
	}
	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Set("apikey", apiKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send request"})
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process resume"})
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read response"})
	}

	var parsedData map[string]interface{}
	err = json.Unmarshal(responseBody, &parsedData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse response"})
	}

	resumeFileAddress := fmt.Sprintf("/resumes/%s/%s", userID, file.Filename)
	err = c.SaveFile(file, resumeFileAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
	}

	profile := models.Profile{
		Applicant:         models.User{ID: userID},
		ResumeFileAddress: resumeFileAddress,
		Skills:            parsedData["skills"].(string),
		Education:         parsedData["education"].(string),
		Experience:        parsedData["experience"].(string),
		Name:              parsedData["name"].(string),
		Email:             parsedData["email"].(string),
		Phone:             parsedData["phone"].(string),
	}

	collection := database.DB.Collection("profiles")

	_, err = collection.InsertOne(context.Background(), profile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save profile"})
	}

	return c.JSON(fiber.Map{"message": "Resume uploaded and processed successfully"})
}
