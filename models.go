package main

type Reminder struct {
	Day               string
	Notification_time string
	Repeat            bool
	Message           string
	Token             string
}

type User struct {
	Name     string
	Password string
	Email    string
}
