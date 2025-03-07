package model

import (
	"time"
)

type Article struct {
	ID         string       `db:"id"`
	Title      string       `db:"title"`
	Slug       string    `db:"slug"`
	CategoryID string       `db:"category_id"`
	Body       string       `db:"body"`
	ImagePath  string       `db:"image_path"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
}