package services

import (
	"go-jwt/internal/database"
	"go-jwt/models"
	"log"
	"time"

	"github.com/google/uuid"
)

func CreateRefreshToken(user models.User) (models.RefreshToken, error) {
	query := "Insert into refresh_token (user_id, token, expiry_date, active) SELECT $1, $2, $3, true WHERE not exists (Select 1 from refresh_token where user_id = $1 and active = true ) RETURNING *"
	stmt, queryErr := database.DB.Prepare(query)

	if queryErr != nil {
		log.Printf("Failed to prepare statement: %v", queryErr)
		return models.RefreshToken{}, queryErr
	}

	refreshToken := models.RefreshToken{
		User_id:     user.Id,
		Token:       uuid.NewString(),
		Expiry_date: time.Now().Add(time.Duration(1000)).Format("2006-01-02 15:04:05"),
		Active:      true,
	}

	insertErr := stmt.QueryRow(refreshToken.User_id, refreshToken.Token, refreshToken.Expiry_date)

	if insertErr.Err() != nil {
		log.Printf("Failed to insert refreshToken: %v", insertErr.Err())
		return models.RefreshToken{}, insertErr.Err()
	}

	return refreshToken, nil
}

func InvalidateTokenForUser(user models.User) {

	query := "Update refresh_token set active = false where user_id = $1"
	stmt, queryErr := database.DB.Prepare(query)

	if queryErr != nil {
		log.Printf("Failed to prepare statement: %v", queryErr)
	}

	row := stmt.QueryRow(user.Id)

	if row.Err() != nil {
		log.Printf("Failed to delete refresh token for user_id: %v", user.Id)
	}
}
