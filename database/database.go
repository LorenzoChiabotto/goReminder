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

func Register_user(name string, password string, email string) error {
	num, err := strconv.Atoi((os.Getenv("dbPort")))
	if err != nil {
		log.Fatal("Error parsing port", err)
	}

	query := url.Values{}
	query.Add("database", os.Getenv("db"))
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(os.Getenv("dbUser"), os.Getenv("dbPassword")),
		Host:     fmt.Sprintf("%s:%d", os.Getenv("dbServer"), num),
		RawQuery: query.Encode(),
	}
	db, err := gorm.Open(sqlserver.Open(u.String()), &gorm.Config{})
	if err != nil {
		return err
	}

	user := Users{
		Name:     name,
		Password: password,
		Email:    email,
		Secret:   "This is a Placeholder",
	}

	result := db.Create(&user)

	return result.Error
}
