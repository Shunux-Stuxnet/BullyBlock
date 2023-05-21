package server

import (
	"EPICS/chat"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Serve() {

	app := fiber.New()
	app.Static("/static", "./static")
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Minute * 5,
	})

	chat.LoadIntentsFromFile("intents.json")
	app.Get("/query", func(c *fiber.Ctx) error {
		return c.SendFile("views/chat.html")
	})

	app.Get("/chat", chat.ChatHandler)
	app.Post("/chat", chat.ChatHandler)
	app.Get("/user", ShowUserDataForm)
	app.Post("/user", ShowUserData)

	app.Post("/showData", func(c *fiber.Ctx) error {
		id := c.FormValue("id")
		user, complain, err := RetrieveData(DB, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
		}
		data := fiber.Map{
			"User":     user,
			"Complain": complain,
		}

		return c.Render("/views/user.html", data)
	})

	fmt.Println("Chatbot is ready. Start typing your messages.")

	app.Get("/", Index)
	app.Get("/about", About)
	app.Get("/blog", Blog)

	app.Use(AuthMiddeware)
	app.Get("/form", Report)
	app.Post("/form", ComplaintData)
	app.Get("/user", Login)
	app.Get("/google_login", GoogleLogin)
	app.Get("/google_callback", GoogleCallback)

	log.Fatal(app.Listen(":8080"))
}
