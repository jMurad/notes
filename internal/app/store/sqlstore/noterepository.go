package sqlstore

import (
	"context"
	"fmt"
	"time"

	"github.com/jMurad/notes/internal/app/model"
	"github.com/jMurad/notes/internal/app/store"
)

type NoteRepository struct {
	store *Store
}

// Create ...
func (r *NoteRepository) Create(ctx context.Context, n *model.Note) error {
	if err := n.Validate(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return r.store.db.QueryRowContext(
		ctx,
		"INSERT INTO notes (content, created, author_id) VALUES ($1, $2, $3) RETURNING id",
		n.Content,
		n.Created,
		n.AuthorID,
	).Scan(&n.ID)
}

// GetNotes ...
func (r *NoteRepository) GetNotes(ctx context.Context, authorId int) (*[]model.Note, error) {
	n := new(model.Note)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	rows, err := r.store.db.QueryContext(
		ctx,
		"SELECT id, content, created, author_id FROM notes WHERE author_id = $1",
		authorId,
	)
	if err != nil {
		return nil, fmt.Errorf("QueryContext: %v", err)
	}
	defer rows.Close()

	nSlice := new([]model.Note)
	for rows.Next() {
		if err = rows.Scan(
			&n.ID,
			&n.Content,
			&n.Created,
			&n.AuthorID,
		); err != nil {
			return nil, err
		}

		*nSlice = append(*nSlice, *n)
	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(*nSlice) == 0 {
		return nil, store.ErrRecordNotFound
	}

	return nSlice, nil
}
