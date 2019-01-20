package controllers

import (
  "net/http"

  beer "data/models"

  "github.com/go-chi/chi"
  "github.com/go-chi/render"
)

func Routes() *chi.Mux {
  router := chi.NewRouter()

  router.Get("/beers", Index)
  router.Get("/beers/{id}", Show)

  return router
}

func Index(w http.ResponseWriter, r *http.Request) {
  result := beer.GetBeers()

  render.JSON(w, r, result)
}

func Show(w http.ResponseWriter, r *http.Request) {
  beer_id := chi.URLParam(r, "id")

  result := beer.GetBeerById(beer_id)

  render.JSON(w, r, result)
}