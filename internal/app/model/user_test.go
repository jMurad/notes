package model_test

import (
	"testing"

	"github.com/jMurad/notes/internal/app/model"
)

func TestUser_CreateEncPass(t *testing.T) {
	u := model.TestUser(t)
	err := u.CreateEncPass()
	if err != nil {
		t.Errorf("Received unexpected error: %v", err)
	}

	if u.EncryptedPassword == "" {
		t.Errorf("Should not be empty, but was %v", u.EncryptedPassword)
	}
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "invalid email and password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				u.Password = ""

				return u

			},
			isValid: false,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "short"

				return u
			},
			isValid: false,
		},
		{
			name: "encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpassword"

				return u
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				if err := tc.u().Validate(); err != nil {
					t.Errorf("Error [%s]: %v", tc.name, err)
				}
			} else {
				if err := tc.u().Validate(); err == nil {
					t.Errorf("Error [%s]: An error is expected but got nil.", tc.name)
				}
			}
		})
	}
}
