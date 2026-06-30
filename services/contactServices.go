package services

import (
	"go-jwt/internal/database"
	"log"
)

type Contact struct {
	Id         uint
	First_Name string
	Last_Name  string
}

func GetContacts() []Contact {
	query := "SELECT * FROM Contacts"
	stmt, queryErr := database.DB.Prepare(query)

	if queryErr != nil {
		log.Printf("Failed to prepare statement: %v", queryErr)
	}

	var contacts []Contact
	rows, querryError := stmt.Query()

	if querryError != nil {
		log.Printf("Failed to retrieve Contacts from the database: %v.", querryError.Error())
		return nil
	}

	for rows.Next() {
		var contact Contact

		if err := rows.Scan(&contact.Id, &contact.First_Name, &contact.Last_Name); err != nil {
			log.Printf("Failed to retrieve the data from the result Query: %v.", err.Error())
			return nil
		}
		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Failed to retrieve the data from the result Query: %v.", err.Error())
		return nil
	}

	return contacts
}
