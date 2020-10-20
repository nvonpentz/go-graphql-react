package middleware

import (
  "context"
  "net/http"

  models "github.com/nvonpentz/go-graphql-react/internal/models/db"
  "github.com/nvonpentz/go-graphql-react/internal/services"
)

// Session handlers gets the user from the session and adds it to the context
func SessionHandler(service *services.Service) func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler { // I wonder if I can just delete this line and init 'next' on the outer call
    // Return this func as HandlerFunc which implements the http.Handler interface
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      session, err := service.SessionStore.Get(r, "go-graphql-react-session")
      if err != nil {
        panic(err)
      }

      var ctx context.Context

      // If session has a user get the user and put it in the context
      userID := session.Values["userID"]
      userIDString, ok := userID.(string)
      if ok {
        user, err := models.FindUser(r.Context(), service.Postgres, userIDString)
        if err != nil {
          panic(err)
        }
        ctx = context.WithValue(r.Context(), "user", user)
      }

      // Add the request and response in the context so we can set headers in
      // resolvers if we choose
      ctx = context.WithValue(r.Context(), "response", w)
      ctx = context.WithValue(ctx, "request", r)

      r = r.WithContext(ctx)
      next.ServeHTTP(w, r)
    })
  }
}
