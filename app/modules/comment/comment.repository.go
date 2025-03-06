package comment

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type CommentRepository struct {
	db *goqu.Database
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: goqu.New("postgres", db)}
}

func (r *CommentRepository) CreateComment(comment *Comment) error {
	query, _, _ := goqu.Insert("comments").Rows(comment).ToSQL()
	_, err := r.db.Exec(query)
	return err
}

func (r *CommentRepository) GetCommentByID(id string) (*Comment, error) {
	var comment Comment
	query := r.db.From("comments").Where(goqu.C("id").Eq(id))

	found, err := query.ScanStruct(&comment)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, sql.ErrNoRows
	}
	return &comment, nil
}

func (r *CommentRepository) UpdateComment(id string, body string) error {
	query, _, _ := goqu.Update("comments").Set(goqu.Record{"body": body, "updated_at": goqu.L("NOW()")}).Where(goqu.C("id").Eq(id)).ToSQL()
	_, err := r.db.Exec(query)
	return err
}

func (r *CommentRepository) DeleteCommentWithReplies(id string) error {
	queryReplies, _, _ := goqu.Delete("comments").Where(goqu.C("parent_id").Eq(id)).ToSQL()
	_, err := r.db.Exec(queryReplies)
	if err != nil {
		return err
	}

	query, _, _ := goqu.Delete("comments").Where(goqu.C("id").Eq(id)).ToSQL()
	_, err = r.db.Exec(query)
	return err
}
