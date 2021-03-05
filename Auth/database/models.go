package database

import mssql "github.com/denisenkom/go-mssqldb"

//Users defines the model in the database
type Users struct {
	ID       mssql.UniqueIdentifier `gorm:"->"`
	Name     string
	Email    string
	Secret   string
	Password string
}
