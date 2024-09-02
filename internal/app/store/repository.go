package store

import (
	"context"

	"github.com/jMurad/notes/internal/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(context.Context, *model.User) error
	Find(context.Context, int) (*model.User, error)
	FindByEmail(context.Context, string) (*model.User, error)
}

// NoteRepository ...
type NoteRepository interface {
	Create(context.Context, *model.Note) error
	GetNotes(context.Context, int) (*[]model.Note, error)
}
