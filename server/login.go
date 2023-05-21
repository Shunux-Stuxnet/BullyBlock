package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func GoogleLogin(c *fiber.Ctx) error {

	url := AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)

}

var uid string

func GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	googlecon := GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}
	var user User
	if err := json.Unmarshal(userData, &user); err != nil {
		log.Fatalf("%s", err.Error())
	}
	InsertUser(user)
	sess, err := store.Get(c)
	if err != nil {
		return c.SendString("Error in handeling the session")
	}
	sess.Set(AUTH_KEY, true)
	sess.Set(USER_ID, user.ID)
	uuid := User{ID: user.ID}
	uid = uuid.ID
	if err := sess.Save(); err != nil {
		return c.SendString("Unable to set session for the user ")
	}
	return c.Redirect("/form")
}

func Index(c *fiber.Ctx) error {
	return c.Render("views/index.html", fiber.Map{})

}

func Report(c *fiber.Ctx) error {
	return c.Render("views/form.html", fiber.Map{})

}

func ComplaintData(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	name := form.Value["name"][0]
	age := form.Value["age"][0]
	email := form.Value["email"][0]
	phone := form.Value["phone"][0]
	description := form.Value["description"][0]

	file, err := form.File["image"][0].Open()
	if err != nil {
		return err
	}
	defer file.Close()
	filename := fmt.Sprintf("%s.%s", uid, filepath.Ext(form.File["image"][0].Filename))
	// Check if the file already exists.
	if _, err := os.Stat(filename); err == nil {
		// The file already exists, so append a number to the file name.
		for i := 1; ; i++ {
			filename = fmt.Sprintf("%s(%d).%s", uid, i, filepath.Ext(form.File["image"][0].Filename))

			// Check if the file exists.
			if _, err := os.Stat(filename); err != nil {
				// The file does not exist, so break out of the loop.
				break
			}
		}
	}
	dst, err := os.Create(fmt.Sprintf("uploads/%s", filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	user := Complain{ID: uid, Name: name, Age: age, Email: email, Phone: phone, Description: description, Image: filename}
	DB.Create(&user)
	fmt.Println(user)
	k := fiber.Map{
		"Name":  user.Name,
		"uuid":  uid,
		"Desc":  user.Description,
		"Image": user.Image,
	}
	fmt.Println(k)
	return c.Render("views/uuid.html", k)
}
