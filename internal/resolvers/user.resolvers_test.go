package resolvers

import (
  "context"
  "fmt"
  "net/http"
  "net/http/httptest"

  models "github.com/nvonpentz/go-graphql-react/internal/models/db"
  graphql "github.com/nvonpentz/go-graphql-react/internal/models/graphql"
)

func (suite *TestSuite) TestSignUp() {
  ctx := context.Background()
  // Add into the session a fake a request and response to simulate
  ctx = context.WithValue(ctx, "response", &httptest.ResponseRecorder{})
  ctx = context.WithValue(ctx, "request", &http.Request{})

  input := graphql.UserInput{
    Email:    "hello@example.com",
    Password: "password",
  }
  mutationResolver := suite.Resolver.Mutation()
  _, err := mutationResolver.SignUp(ctx, input)

  suite.Require().NoError(err, "Failed to resolve SignUp")

  count, err := models.Users().Count(ctx, suite.Resolver.Service.Postgres)
  suite.Require().NoError(err, "Failed to count users")
  suite.Require().Equal(count, int64(1))
}

func (suite *TestSuite) TestLogIn() {
  // Add into the session a fake a request and response to simulate
  ctx := context.Background()
  ctx = context.WithValue(ctx, "response", &httptest.ResponseRecorder{})
  ctx = context.WithValue(ctx, "request", &http.Request{})

  // First create a user via a sign up
  input := graphql.UserInput{
    Email:    "hello@example.com",
    Password: "password",
  }
  mutationResolver := suite.Resolver.Mutation()
  _, err := mutationResolver.SignUp(ctx, input)
  suite.Require().NoError(err, "Failed to resolve SignUp")

  user, err := mutationResolver.LogIn(ctx, input)
  suite.Require().NoError(err, "Failed to resolve LogIn")
  fmt.Println(user)
  suite.Require().Equal(user.Email, input.Email)
}
