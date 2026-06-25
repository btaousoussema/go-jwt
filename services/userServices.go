package services

import (
	"go-jwt/internal/database"
	"go-jwt/models"
	"log"
)

func InsertUser(user models.User) error {
	query := "Insert into users (email, password) SELECT $1, $2 WHERE not exists (Select 1 from users where email = $1 ) RETURNING id"
	stmt, queryErr := database.DB.Prepare(query)

	if queryErr != nil {
		log.Printf("Failed to prepare statement: %v", queryErr)
		return queryErr
	}

	insertErr := stmt.QueryRow(user.Email, user.Password).Scan(&user.Id)

	if insertErr != nil {
		log.Printf("Failed to insert user: %v", insertErr.Error())
		return queryErr
	}

	return nil
}

func GetUser(email string) (models.User, error) {
	query := "SELECT id, email, password from users where email = $1"
	stmt, queryErr := database.DB.Prepare(query)
	if queryErr != nil {
		log.Printf("Failed to prepare statement: %v", queryErr)
		return models.User{}, queryErr
	}

	var user models.User
	err := stmt.QueryRow(email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
