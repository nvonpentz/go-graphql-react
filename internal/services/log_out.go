package services

import (
  "context"
  "net/http"
)

func (service Service) LogOut(ctx context.Context) (*bool, error) {
  // 1. Get the request from the context
  request := ctx.Value("request").(*http.Request)
  response := ctx.Value("response").(http.ResponseWriter)

  // 2. Get the session from the request
  session, err := service.SessionStore.Get(request, "go-graphql-react-session")
  if err != nil {
    newFalse := false
    return &newFalse, err
  }

  // 3. Log out the user and save the new session
  session.Values["userID"] = nil
  session.Values["loggedIn"] = false
  session.Save(request, response)

  newTrue := true
  return &newTrue, nil
}
