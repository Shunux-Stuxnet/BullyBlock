package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Intent struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
	Response string   `json:"response"`
}

type IntentData struct {
	Intents []Intent `json:"intents"`
}

var intents []Intent

func ChatHandler(c *fiber.Ctx) error {
	message := c.Query("message")
	fmt.Println(message)
	if message == "exit" {
		return c.SendString("Exiting...")
	}

	intent := classifyIntent(message)
	response := handleIntent(intent)

	return c.SendString(response)
}

func LoadIntentsFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open intents file:", err)
	}
	defer file.Close()

	var intentData IntentData
	err = json.NewDecoder(file).Decode(&intentData)
	if err != nil {
		log.Fatal("Failed to parse intents file:", err)
	}

	intents = intentData.Intents
}

func classifyIntent(message string) string {
	for _, intent := range intents {
		for _, keyword := range intent.Keywords {
			if strings.Contains(strings.ToLower(message), strings.ToLower(keyword)) {
				return intent.Name
			}
		}
	}

	return "unknown"
}

func handleIntent(intentName string) string {
	for _, intent := range intents {
		if intent.Name == intentName {
			return intent.Response
		}
	}

	return "Sorry, I couldn't understand your request."
}
