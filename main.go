package main

import (
  "fmt"
  "log"
  "net/http"

  beer "controllers"

  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/render"
)


func main() {
  fmt.Println("Login and serve at :4000")
  router := Routes()

  log.Fatal(http.ListenAndServe(":4000", router))
}

func Routes() *chi.Mux {
  router := chi.NewRouter()

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

  return router
}