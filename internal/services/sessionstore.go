package services

import (
  "github.com/gorilla/sessions"
  "os"
)

func NewSessionStoreWithDefaults() *sessions.CookieStore {
  return sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}
