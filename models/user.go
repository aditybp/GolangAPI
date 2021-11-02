package models

type User struct {
	Id uint
	Nama string
	Email string `gorm:"unique"`
	Password []byte
}

