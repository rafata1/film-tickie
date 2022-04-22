package auth

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/templateOfService/connectors/mysql"
)

type Repo struct {
	conn *sqlx.DB
}

func NewRepo() *Repo {
	return &Repo{
		conn: mysql.GetMySQLInstance(),
	}
}

var updateUserInfoQuery = "INSERT INTO users (phone_number, name) VALUE(?, ?)"

func (r *Repo) UpdateUserInfo(phone string, name string) error {
	_, err := r.conn.Exec(updateUserInfoQuery, phone, name)
	if err != nil {
		fmt.Println("error during saving user info")
	}
	return err
}
