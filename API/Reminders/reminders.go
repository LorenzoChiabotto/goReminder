package Reminders

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"net/url"
	"os"
	"strconv"
)

func connect_database() (*gorm.DB, error) {
	num, err := strconv.Atoi((os.Getenv("dbPort")))
	if err != nil {
		log.Fatalf("Error parsing port: %v", err)
	}

	query := url.Values{}
	query.Add("database", os.Getenv("db"))
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(os.Getenv("dbUser"), os.Getenv("dbPassword")),
		Host:     fmt.Sprintf("%s:%d", os.Getenv("dbServer"), num),
		RawQuery: query.Encode(),
	}
	return gorm.Open(sqlserver.Open(u.String()), &gorm.Config{})
}

func getUser(id string) (*Users, error) {
	db, err := connect_database()
	if err != nil {
		return nil, err
	}
	var user Users

	result := db.Find(&user, "ID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Invalid credentials")
	}

	return &user, nil
}

func AddReminder(reminder Reminder, token *jwt.Token) error {
	db, err := connect_database()
	if err != nil {
		return err
	}

	user, err := getUser(token.Claims.(jwt.MapClaims)["userId"].(string))
	if err != nil {
		return err
	}

	result := db.Create(
		&Reminders{
			Users:         	user,
			ReminderTime:   reminder.Day,
			Repeat:         reminder.Repeat,
			Email: user.Email,
			MessageTitle: 	"GoReminder",
			MessageBody:    reminder.Message,
	})

	return result.Error
}