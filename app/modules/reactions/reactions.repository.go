package reactions

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type ReactionRepository struct {
	db *goqu.Database
}

func NewReactionRepository(db *sql.DB) *ReactionRepository {
	return &ReactionRepository{db: goqu.New("postgres", db)}
}

func (r *ReactionRepository) CreateReaction(reaction *Reaction) error {
	_, err := r.db.Insert("reactions").Rows(reaction).Executor().Exec()
	return err
}

func (r *ReactionRepository) DeleteReaction(articleID, userID uuid.UUID) error {
	_, err := r.db.Delete("reactions").
		Where(goqu.Ex{"article_id": articleID, "user_id": userID}).
		Executor().Exec()
	return err
}

func (r *ReactionRepository) GetReaction(articleID, userID uuid.UUID) (*Reaction, error) {
	var reaction Reaction
	found, err := r.db.From("reactions").
		Where(goqu.Ex{"article_id": articleID, "user_id": userID}).
		ScanStruct(&reaction)
	if err != nil || !found {
		return nil, err
	}
	return &reaction, nil
}
