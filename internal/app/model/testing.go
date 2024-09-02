package model

import (
	"testing"
)

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

// TestNote ...
func TestNote(t *testing.T) *Note {
	return &Note{
		Content: "Он оперся на заступ и посмотрел на небо",
	}
}
