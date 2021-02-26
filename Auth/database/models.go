package database

//Users defines the model in the database
type Users struct {
	ID       string `gorm:"->"`
	Name     string
	Email    string
	Secret   string
	Password string
}
