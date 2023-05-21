package server

type User struct {
	ID    string `gorm:"primaryKey" json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Complain struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Name        string
	Age         string
	Email       string
	Phone       string
	Description string
	Image       string
}
