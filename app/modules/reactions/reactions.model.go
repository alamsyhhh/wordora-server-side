package reactions

import (
	"time"

	"github.com/google/uuid"
)

type Reaction struct {
	ID        uuid.UUID `db:"id"`
	ArticleID uuid.UUID `db:"article_id"`
	UserID    uuid.UUID `db:"user_id"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
