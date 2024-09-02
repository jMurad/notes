package notes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/jMurad/notes/internal/app/model"
	"github.com/jMurad/notes/internal/app/store/teststore"
	"github.com/sirupsen/logrus"
)

const (
	logLevel = logrus.ErrorLevel
)

func TestServer_AuthenticateUser(t *testing.T) {
	u := model.TestUser(t)

	store := teststore.New()
	store.User().Create(context.Background(), u)

	secretKey := []byte("secret")
	s := newServer(store, sessions.NewCookieStore(secretKey))
	s.logger.Level = logLevel

	sc := securecookie.New(secretKey, nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticate",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticate",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			if tc.expectedCode != rec.Code {
				t.Errorf("Not equal:\nexpected: %d\nactual  : %d", tc.expectedCode, rec.Code)
			}
		})
	}
}

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	s.logger.Level = logLevel

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
				// "password": "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			if tc.expectedCode != rec.Code {
				t.Errorf("Not equal:\nexpected: %d\nactual  : %d", tc.expectedCode, rec.Code)
			}
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)

	store := teststore.New()
	store.User().Create(context.Background(), u)

	s := newServer(store, sessions.NewCookieStore([]byte("secret")))
	s.logger.Level = logLevel

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			if tc.expectedCode != rec.Code {
				t.Errorf("Not equal:\nexpected: %d\nactual  : %d", tc.expectedCode, rec.Code)
			}
		})
	}
}

func TestServer_HandleNotesCreate(t *testing.T) {
	u := model.TestUser(t)
	n := model.TestNote(t)

	store := teststore.New()
	store.User().Create(context.Background(), u)

	secretKey := []byte("secret")
	s := newServer(store, sessions.NewCookieStore(secretKey))
	s.logger.Level = logLevel

	sc := securecookie.New(secretKey, nil)

	cookieValue := map[interface{}]interface{}{
		"user_id": u.ID,
	}

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"content": n.Content,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid content",
			payload: map[string]interface{}{
				"content": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			bb := &bytes.Buffer{}
			json.NewEncoder(bb).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/private/notes", bb)
			cookieStr, _ := sc.Encode(sessionName, cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.ServeHTTP(rec, req)

			if tc.expectedCode != rec.Code {
				t.Errorf("Not equal:\nexpected: %d\nactual  : %d", tc.expectedCode, rec.Code)
			}
		})
	}
}

func TestServer_HandleNotesGet(t *testing.T) {
	u := model.TestUser(t)

	store := teststore.New()
	store.User().Create(context.Background(), u)

	secretKey := []byte("secret")
	s := newServer(store, sessions.NewCookieStore(secretKey))
	s.logger.Level = logLevel

	sc := securecookie.New(secretKey, nil)

	cookieValue := map[interface{}]interface{}{
		"user_id": u.ID,
	}

	req, _ := http.NewRequest(http.MethodGet, "/private/notes", nil)
	cookieStr, _ := sc.Encode(sessionName, cookieValue)
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	if http.StatusNoContent != rec.Code {
		t.Errorf("Not equal:\nexpected: %d\nactual  : %d", http.StatusNoContent, rec.Code)
	}

	n := model.TestNote(t)
	store.Note().Create(context.Background(), n)

	rec = httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	if http.StatusOK != rec.Code {
		t.Errorf("Not equal:\nexpected: %d\nactual  : %d", http.StatusOK, rec.Code)
	}
}
