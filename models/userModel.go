package model 

type User struct {
	Id uint `gorm:"primaryKey"`
	Email string `gorm:"unique"`
	Password string
}