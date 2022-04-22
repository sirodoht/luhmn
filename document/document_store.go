package document

import (
	"context"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -destination=document_store_mock_generated.go -package=document -source=document_store.go
type Store interface {
	Insert(context.Context, *Document) (int64, error)
	GetAll(context.Context) ([]*Document, error)
}

type SQLStore struct {
	db *sqlx.DB
}

func NewSQLStore(db *sqlx.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (s *SQLStore) Insert(ctx context.Context, d *Document) (int64, error) {
	res, err := s.db.NamedExec(`
		INSERT INTO documents (
			title,
			body,
			created_at,
			updated_at
		) VALUES (
			:title,
			:body,
			:created_at,
			:updated_at
		)`, d)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SQLStore) GetAll(ctx context.Context) ([]*Document, error) {
	var docs []*Document
	err := s.db.SelectContext(
		ctx,
		&docs,
		`SELECT * FROM documents ORDER BY id ASC`,
	)
	if err != nil {
		return nil, err
	}
	return docs, nil
}
