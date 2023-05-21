package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func FetchUserDataByID(id string) (User, Complain, error) {
	var user User
	var complain Complain

	err := DB.Model(&User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return User{}, Complain{}, fmt.Errorf("failed to fetch user data: %v", err)
	}

	err = DB.Model(&Complain{}).Where("id = ?", id).First(&complain).Error
	if err != nil {
		return User{}, Complain{}, fmt.Errorf("failed to fetch complain data: %v", err)
	}

	return user, complain, nil
}

func ShowUserDataForm(c *fiber.Ctx) error {
	return c.Render("views/user_form.html", fiber.Map{})
}

func ShowUserData(c *fiber.Ctx) error {
	id := c.FormValue("id")
	user, complain, err := FetchUserDataByID(id)
	if err != nil {
		log.Printf("Failed to fetch user data: %v", err)
		return c.SendString("Failed to fetch user data")
	}

	return c.Render("views/user.html", fiber.Map{
		"User":     user,
		"Complain": complain,
	})
}
