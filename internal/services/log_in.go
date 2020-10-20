package services

import (
  "context"
  "net/http"

  models "github.com/nvonpentz/go-graphql-react/internal/models/db"
  graphql "github.com/nvonpentz/go-graphql-react/internal/models/graphql"
  "golang.org/x/crypto/bcrypt"
)

func (service Service) LogIn(ctx context.Context, input graphql.UserInput) (*models.User, error) {
  user, err := models.Users(models.UserWhere.Email.EQ(input.Email)).One(ctx, service.Postgres)
  if err != nil {
    return &models.User{}, nil
  }

  passwordMatches := checkPasswordHash(input.Password, user.Password)

  if !passwordMatches {
    return &models.User{}, nil
  }

  // Set the session
  // 1. Get the request from the context
  request := ctx.Value("request").(*http.Request)
  response := ctx.Value("response").(http.ResponseWriter)

  // 2. Get the session from the request
  session, err := service.SessionStore.Get(request, "go-graphql-react-session")
  if err != nil {
    return user, err
  }

  // 3. Set userID and loggedIn to true
  session.Values["userID"] = user.ID
  session.Values["loggedIn"] = true
  session.Save(request, response)

  return user, nil
}

func checkPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}
