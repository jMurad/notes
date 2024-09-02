package model_test

import (
	"testing"

	"github.com/jMurad/notes/internal/app/model"
)

func TestNote_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		n       func() *model.Note
		isValid bool
	}{
		{
			name: "valid",
			n: func() *model.Note {
				return model.TestNote(t)
			},
			isValid: true,
		},
		{
			name: "empty content",
			n: func() *model.Note {
				n := model.TestNote(t)
				n.Content = ""

				return n
			},
			isValid: false,
		},
		{
			name: "content with errors",
			n: func() *model.Note {
				n := model.TestNote(t)
				n.Content = "Летам мы чясто ходим в лес за грибами и ягадами"

				return n
			},
			isValid: false,
		},
		{
			name: "short content",
			n: func() *model.Note {
				n := model.TestNote(t)
				n.Content = "Сутра марасит"

				return n
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				if err := tc.n().Validate(); err != nil {
					t.Errorf("Error [%s]: %v", tc.name, err)
				}
			} else {
				if err := tc.n().Validate(); err == nil {
					t.Errorf("Error [%s]: An error is expected but got nil.", tc.name)
				}
			}
		})
	}
}
