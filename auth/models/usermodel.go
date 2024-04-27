package models

import (
	"database/sql"
	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}
	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {

	row, err := u.db.Query("SELECT username,nama, password FROM kasir WHERE "+fieldName+"= ? limit 1", fieldValue)

	if err != nil {
		return err
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&user.Username, &user.Nama, &user.Password)
	}
	return nil
}
