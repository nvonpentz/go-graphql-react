package services

import "github.com/gorilla/sessions"

type Service struct {
	Postgres     *Postgres
	SessionStore *sessions.CookieStore
}

func NewWithDefaults() (*Service, error) {
	postgres, err := NewPostgresWithDefaults()
	if err != nil {
		return &Service{}, err
	}
	sessionStore := NewSessionStoreWithDefaults()

	return &Service{
		Postgres:     postgres,
		SessionStore: sessionStore,
	}, nil
}
