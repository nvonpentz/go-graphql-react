package internal

import (
  "net/http"

  "github.com/99designs/gqlgen/graphql/handler"
  "github.com/99designs/gqlgen/graphql/playground"
  "github.com/go-chi/chi"
  chiware "github.com/go-chi/chi/middleware"
  "github.com/go-chi/cors"
  "github.com/nvonpentz/go-graphql-react/internal/middleware"
  "github.com/nvonpentz/go-graphql-react/internal/resolvers"
  "github.com/nvonpentz/go-graphql-react/internal/services"
)

func NewRouterWithDefaults() (chi.Router, error) {
  service, err := services.NewWithDefaults()
  if err != nil {
    return nil, err
  }

  router := chi.NewRouter()
  router.Use(chiware.Logger)
  router.Use(middleware.SessionHandler(service))
  // FIXME
  router.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"*"},
  }))

  // Frontend routes
  reactFilesServer := http.FileServer(http.Dir("./frontend/build"))
  router.Method("GET", "/*", reactFilesServer)
  router.Method("GET", "/sign_up", http.StripPrefix("/sign_up", reactFilesServer))
  router.Method("GET", "/log_in", http.StripPrefix("/log_in", reactFilesServer))

  // Backend routes
  resolver := &resolvers.Resolver{
    Service: service,
  }
  if err != nil {
    return nil, err
  }
  graphQLServer := handler.NewDefaultServer(
    resolvers.NewExecutableSchema(
      resolvers.Config{
        Resolvers: resolver,
      },
    ),
  )
  router.Mount("/playground", playground.Handler("GraphQL playground", "/graphql"))
  router.Mount("/graphql", graphQLServer)

  return router, nil
}
