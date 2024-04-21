package models

import (
	"database/sql"
	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
)

type Usermodel struct {
	db *sql.DB
}
func NewUserModel() *Usermodel{
	conn, err := config.DBConnect()

	if err !=nil{
		panic(err)
	}
	return &Usermodel{
		db: conn,
	}
}

func (u Usermodel) Where(user *entities.User, fieldname, fieldValue string) error  {

	row, err := u.db.Query("select name, password from member where" + fieldname + " = ? limit 1", fieldValue)
	if err != nil {
		return err
	}
	defer row.Close()
	for row.Next(
		row.Scan(&user.id))
}	return nil