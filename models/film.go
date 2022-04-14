package models

import "time"

type Film struct {
    Id          int       `db:"id"`
    Name        string    `db:"name"`
    Description string    `db:"description"`
    Length      int       `db:"length"`
    OpeningDay  time.Time `db:"opening_day"`
    Category    string    `db:"category"`
    ImageUrls   string    `db:"image_urls"`
    CreatedAt   time.Time `db:"created_at"`
    UpdatedAt   time.Time `db:"updated_at"`
}
