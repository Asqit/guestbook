package models

import (
	"errors"
	"log"
	"time"

	"github.com/asqit/guestbook/internal/db"
)

// Structure is inspired by: http://maplecity.com/homepage/guest1.html

type Guest struct {
	ID      *int
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Comment string    `json:"comment"`
	Time    time.Time `json:"time"`
}

func CreateGuestTable() error {
	if db.DB == nil {
		return errors.New("failed to create a \"guests\" table, db connector is nil")
	}

	query := `
		CREATE TABLE IF NOT EXISTS guests (
			ID SERIAL PRIMARY KEY,
			time TIMESTAMP NOT NULL,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			comment TEXT
		);
	`

	_, err := db.DB.Exec(query)
	return err
}

func InsertNewGuest(guest *Guest) error {
	if db.DB == nil {
		return errors.New("failed to connect to the \"guests\" table, db connector is nil")
	}
	query := `
		INSERT INTO guests (name, email, comment, time) VALUES ($1, $2, $3, $4);
	`

	_, err := db.DB.Exec(query, guest.Name, guest.Email, guest.Comment, time.Now())

	return err
}

// A function that will return at most 50 values starting from
// desired ID parameter.
func GetAllGuests() ([]Guest, error) {
	if db.DB == nil {
		return nil, errors.New("failed to connect to the \"guests\" table, db connector is nil")
	}

	query := `SELECT * FROM guests;`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	slice := []Guest{}
	for rows.Next() {
		var guest Guest
		err := rows.Scan(&guest.ID, &guest.Time, &guest.Name, &guest.Email, &guest.Comment)
		if err != nil {
			log.Fatal(err)
		}

		slice = append(slice, guest)
	}

	return slice, nil
}
