package teststore

import (
	"context"

	"github.com/jMurad/notes/internal/app/model"
	"github.com/jMurad/notes/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// Create ...
func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.CreateEncPass(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// Find ...
func (r *UserRepository) Find(ctx context.Context, id int) (*model.User, error) {
	if u, ok := r.users[id]; ok {
		return u, nil
	} else {
		return nil, store.ErrRecordNotFound
	}
}
