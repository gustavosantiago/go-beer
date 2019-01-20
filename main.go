package main

import (
  "fmt"
  "log"
  "net/http"

  "data/gql"
  "controllers/server"
  beer "controllers"

  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/render"
  "github.com/graphql-go/graphql"
)


func main() {
  fmt.Println("Login and serve at :4000")

  router := initializeAPI()

  log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() *chi.Mux {
  router := chi.NewRouter()

  rootQuery := gql.NewRoot()

  source, err := graphql.NewSchema(
    graphql.SchemaConfig{Query: rootQuery.Query},
  )

  if err != nil {
    fmt.Println("Error creating schema: ", err)
  }

  // Create a server struct that holds a poiter to db
  s := server.Setup{
    GqlSchema: &source,
  }

  router.Use(
    render.SetContentType(render.ContentTypeJSON),
    middleware.Logger,
    middleware.DefaultCompress,
    middleware.RedirectSlashes,
    middleware.Recoverer,
  )

  router.Route("/", func(r chi.Router) {
    r.Mount("/", beer.Routes())
  })

  router.Post("/graphql", s.GraphQL())

  return router
}
