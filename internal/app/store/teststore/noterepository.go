package teststore

import (
	"context"

	"github.com/jMurad/notes/internal/app/model"
	"github.com/jMurad/notes/internal/app/store"
)

// NoteRepository ...
type NoteRepository struct {
	store *Store
	notes map[int]*model.Note
}

// Create ...
func (r *NoteRepository) Create(ctx context.Context, n *model.Note) error {
	if err := n.Validate(); err != nil {
		return err
	}

	n.ID = len(r.notes) + 1
	r.notes[n.ID] = n

	return nil
}

// GetNotes ...
func (r *NoteRepository) GetNotes(ctx context.Context, authorId int) (*[]model.Note, error) {
	nSlice := new([]model.Note)
	if len(r.notes) == 0 {
		return nil, store.ErrRecordNotFound
	}

	for _, n := range r.notes {
		if n.AuthorID == authorId {
			*nSlice = append(*nSlice, *n)
		}
	}

	return nSlice, nil
}
