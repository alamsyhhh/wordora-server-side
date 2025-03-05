package comment

import "time"

type Comment struct {
	ID        string    `db:"id"`
	ArticleID string    `db:"article_id"`
	UserID    string    `db:"user_id"`
	ParentID  *string   `db:"parent_id"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
