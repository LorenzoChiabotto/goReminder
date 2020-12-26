package database

type Users struct {
	Id       string `gorm:"->"`
	Name     string
	Email    string
	Secret   string
	Password string
}

type Reminder struct {
	User            Users
	Day             string
	Repeat          bool
	Message_subject string
	Message_body    string
	Token           string
}
