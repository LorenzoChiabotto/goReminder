package main

import "time"

//AddReminderRequest represents a request to create a reminder
type AddReminderRequest struct {
	Day              time.Time
	Repeat           bool
	Message          string
}


//PutReminderRequest represents a request to edit a reminder
type PutReminderRequest struct {
	ID              string
	Day              time.Time
	Repeat           bool
	Message          string
	Token            string
}
//DeleteReminderRequest represents a request to delete a reminder
type DeleteReminderRequest struct {
	ID              string
}
//GetReminderRequest represents a request to get a reminder
type GetReminderRequest struct {
	ID              string
}