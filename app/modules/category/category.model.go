package category

import "time"

type Category struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
