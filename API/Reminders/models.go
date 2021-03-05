package Reminders

import (
	mssql "github.com/denisenkom/go-mssqldb"
	"time"
)


type Users struct {
	Id       mssql.UniqueIdentifier `gorm:"->"`
	Name     string
	Email    string
	Secret   string
	Password string
}

//Reminder represents a reminder
type Reminder struct {
	Day              time.Time
	Repeat           bool
	Message          string
	Token            string
}
//Reminders represents a reminder in the database
type Reminders struct {
	Id 				mssql.UniqueIdentifier `gorm:"->"`
	UserId        	mssql.UniqueIdentifier
	Users          *Users `gorm:"foreignKey:UserId"`
	Email 			string
	ReminderTime	time.Time
	Repeat          bool
	MessageTitle 	string
	MessageBody    	string
}