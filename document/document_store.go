package document

import (
	"fmt"
	"context"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -destination=document_store_mock_generated.go -package=document -source=document_store.go
type Store interface {
	Insert(context.Context, *Document) (int64, error)
	GetAll(context.Context) ([]*Document, error)
	Update(context.Context, int64, string, string) (error)
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

func (s *SQLStore) Update(ctx context.Context, id int64, field string, value string) (error) {
	sql := fmt.Sprintf("UPDATE documents SET %s=:value WHERE id=:id", field)
	_, err := s.db.NamedExec(sql, map[string]interface{}{
			"field": field,
			"value": value,
			"id": id,
		})
	if err != nil {
		return err
	}
	return nil
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

func (s *SQLStore) GetOne(ctx context.Context, id int64) (*Document, error) {
	var doc *Document
	err := s.db.SelectContext(
		ctx,
		&doc,
		`SELECT * FROM documents WHERE id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
