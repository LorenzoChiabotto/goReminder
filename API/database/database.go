package database

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func connectDatabase() (*gorm.DB, error) {
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

func addReminder() (*Reminders, error){

	return nil, nil
}