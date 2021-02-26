package database

//Users defines the model in the database
type Users struct {
	ID       string `gorm:"->"`
	Name     string
	Email    string
	Secret   string
	Password string
}

//Reminders defines the model in the database
type Reminders struct {
	User           Users
	Day            string
	Repeat         bool
	MessageSubject string
	MessageBody    string
	Token          string
}
