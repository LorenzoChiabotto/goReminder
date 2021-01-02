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

func Get_user(name string, password string) (*Users, error) {
	db, err := connect_database()
	if err != nil {
		return nil, err
	}
	var user Users

	result := db.Find(&user, "name = ? AND password = ?", name, password)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Invalid credentials")
	}

	return &user, nil
}

func Get_secret(username string) (*string, error) {
	db, err := connect_database()
	if err != nil {
		return nil, err
	}

	var user Users
	result := db.First(&user, "name = ? ", username)
	if result.Error != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}
	return &(user.Secret), nil
}

func Register_user(name string, password string, email string) error {
	db, err := connect_database()
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
