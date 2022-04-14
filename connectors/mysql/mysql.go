package mysql

import (
    "github.com/jmoiron/sqlx"
    "os"
    "time"
)

var conn *sqlx.DB

func Connect() error {
    var err error
    conn, err = sqlx.Connect("mysql", os.Getenv("MYSQL"))
    conn.SetMaxIdleConns(50)
    conn.SetMaxOpenConns(100)
    conn.SetConnMaxLifetime(10 * time.Second)
    return err
}

func GetMySQLInstance() *sqlx.DB {
    return conn
}
