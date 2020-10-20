package services

import (
  "context"
  "net/http"
  "regexp"

  gqlgen "github.com/99designs/gqlgen/graphql"
  models "github.com/nvonpentz/go-graphql-react/internal/models/db"
  graphql "github.com/nvonpentz/go-graphql-react/internal/models/graphql"
  "github.com/volatiletech/sqlboiler/v4/boil"
  "golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (service Service) SignUp(ctx context.Context, input graphql.UserInput) (*models.User, error) {
  if !isEmailValid(input.Email) {
    // Add errors and return
    gqlgen.AddErrorf(ctx, "%s is not a valid email.", input.Email)
  }

  hashedPassword, err := hashPassword(input.Password)
  if err != nil {
    return &models.User{}, err
  }

  newUser := &models.User{
    Name:     *input.Name,
    Email:    input.Email,
    Password: hashedPassword,
  }

  err = newUser.Insert(ctx, service.Postgres, boil.Infer())

  // 1. Get the request from the context
  request := ctx.Value("request").(*http.Request)
  response := ctx.Value("response").(http.ResponseWriter)

  // 2. Get the session from the request
  session, err := service.SessionStore.Get(request, "go-graphql-react-session")
  if err != nil {
    return newUser, err
  }

  // 3. Set userID and loggedIn to true
  session.Values["userID"] = newUser.ID
  session.Values["loggedIn"] = true
  session.Save(request, response)

  return newUser, err
}

func hashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  return string(bytes), err
}

// https://golangcode.com/validate-an-email-address/
func isEmailValid(email string) bool {
  if len(email) < 3 && len(email) > 254 {
    return false
  }
  return emailRegex.MatchString(email)
}
