package Reminders

//Reminder represents a reminder
type Reminder struct {
	Day              string
	NotificationTime string
	Repeat           bool
	Message          string
	Token            string
}