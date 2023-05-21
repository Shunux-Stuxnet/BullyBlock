package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	dns := "shunux:P@s$w0rd@tcp(127.0.0.1:3306)/epics?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: false})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	err = DB.AutoMigrate(&User{}, &Complain{})
	if err != nil {
		log.Fatal("Failed to migrate database schema")
	}
	fmt.Println("Connected successfully to the database")
}

func InsertUser(data User) {
	val := DB.Create(data)
	fmt.Println(val)
}
func InsertComplain(data Complain) {
	val := DB.Create(data)
	fmt.Println(val)
}

func About(c *fiber.Ctx) error {
	return c.Render("views/about.html", fiber.Map{})
}
func Login(c *fiber.Ctx) error {
	return c.Render("views/user.html", fiber.Map{})
}
func Blog(c *fiber.Ctx) error {
	return c.Render("views/blog.html", fiber.Map{})
}

func RetrieveData(db *gorm.DB, id string) (User, Complain, error) {
	var user User
	var complain Complain

	err := db.First(&user, id).Error
	if err != nil {
		return user, complain, err
	}

	err = db.First(&complain, id).Error
	if err != nil {
		return user, complain, err
	}

	return user, complain, nil
}



