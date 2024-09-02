package sqlstore_test

import (
	"context"
	"testing"

	"github.com/jMurad/notes/internal/app/model"
	"github.com/jMurad/notes/internal/app/store"
	"github.com/jMurad/notes/internal/app/store/sqlstore"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	u := model.TestUser(t)

	if err := s.User().Create(context.Background(), u); err != nil {
		t.Errorf("Received unexpected error: %v", err)
	}
	if u == nil {
		t.Errorf("Expected value not to be nil")
	}
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	u1 := model.TestUser(t)

	_, err := s.User().FindByEmail(context.Background(), u1.Email)
	if err != store.ErrRecordNotFound {
		t.Errorf("Error message not equal:\nexpected: %q\nactual  : %q", store.ErrRecordNotFound, err)
	}

	s.User().Create(context.Background(), u1)

	u2, err := s.User().FindByEmail(context.Background(), u1.Email)
	if err != nil {
		t.Errorf("Received unexpected error: %v", err)
	}
	if u2 == nil {
		t.Errorf("Expected value not to be nil")
	}
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	u1 := model.TestUser(t)

	s.User().Create(context.Background(), u1)

	u2, err := s.User().Find(context.Background(), u1.ID)

	if err != nil {
		t.Errorf("Received unexpected error: %v", err)
	}
	if u2 == nil {
		t.Errorf("Expected value not to be nil")
	}
}
