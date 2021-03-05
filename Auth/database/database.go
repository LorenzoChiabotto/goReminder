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
	num, err := strconv.Atoi(os.Getenv("dbPort"))
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

//GetUser retrieves an user with name a pw
func GetUser(name string, password string) (*Users, error) {
	db, err := connectDatabase()
	if err != nil {
		return nil, err
	}
	var user Users

	result := db.Find(&user, "name = ? AND password = ?", name, password)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("invalid credentials")
	}

	log.Print(&user)

	return &user, nil
}

//GetSecret retrieves the token for a username
func GetSecret(username string) (*string, error) {
	db, err := connectDatabase()
	if err != nil {
		return nil, err
	}

	var user Users
	result := db.First(&user, "name = ? ", username)
	if result.Error != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("invalid credentials")
	}
	return &(user.Secret), nil
}