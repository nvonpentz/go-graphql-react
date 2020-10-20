# go-graphql-react

A template web app with a React / TypeScript frontend and a Go GraphQL backend that uses [SQLBoiler](https://github.com/volatiletech/sqlboiler) alongside [gqlgen](https://github.com/99designs/gqlgen) for a fast, schema driven development style.

## Stack
* Golang backend server
* Postgres database
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) to generate Go ORM models based on database tables
* [gqlgen](https://github.com/99designs/gqlgen) to generate a graphQL API based on the models
* Gorilla [sessions](https://github.com/gorilla/sessions) for authentication (sign up / log in / log out resolvers are implemented)
* React / TypeScript frontend
* Apollo GraphQL client

## Development
SQLBoiler and gqlgen work really well together to generate nearly an entire graphQL API from the database:
1. Add Postgres table definitions to the `/migrations/` folder
1. Generate ORM based on tables in the database `make db`
1. Add a graphQL schema in the `/api/` folder that reflects the generated ORM
1. Generate a graphQL server from the schema that binds to the ORM `make graphql` (bonus: complete 2 & 4 with `make models`)
1. Fill in the graphQL resolvers using the ORM in `/internal/resolvers/`

## Running
1. Create a new file `.env.sh` at the top level directory, copy in `.env.example.sh` and fill in the missing variables
1. Start the React app `make run-client` and start the backend `make run-server` (migrations are run automatically with [golang-migrate](https://github.com/golang-migrate/migrate))
