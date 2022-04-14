package models

import "time"

type Cinema struct {
    Id          int       `db:"id"`
    Name        string    `db:"name"`
    Description string    `db:"description"`
    Address     string    `db:"address"`
    CreatedAt   time.Time `db:"created_at"`
    UpdatedAt   time.Time `db:"updated_at"`
}
