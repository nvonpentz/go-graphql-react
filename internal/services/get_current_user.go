package services

import (
  "context"
  "net/http"

  models "github.com/nvonpentz/go-graphql-react/internal/models/db"
)

func (service Service) GetCurrentUser(ctx context.Context) (*models.User, error) {
  request := ctx.Value("request").(*http.Request)
  session, err := service.SessionStore.Get(request, "go-graphql-react-session")
  if err != nil {
    return nil, err
  }

  // No authenticated cookie supplied
  userID, ok := session.Values["userID"].(string)
  if !ok {
    return nil, nil
  }

  user, err := models.FindUser(ctx, service.Postgres, userID)
  if err != nil {
    return nil, err
  }
  return user, nil
}
