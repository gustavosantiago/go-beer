package models

import (
  "fmt"

  "data/postgres"
)

type Beer struct {
  ID          int      `json:"id"`
  Name        string   `json:"name"`
  Ibu         float64  `json:"ibu"`
  Abv         float64  `json:"abv"`
  Description string   `json:"description"`
  UpdatedAt   string   `json:"updated_at"`
}

func GetBeers() []Beer {
  var beers []Beer

  err := postgres.Client.Model(&beers).Select()

  if err != nil {
    fmt.Println(err)
  }

  return beers
}

func GetBeerById(id string) Beer {
  var beer Beer

  err := postgres.Client.Model(&beer).
    Where("id = ?", id).Limit(1).Select()

  if err != nil {
    fmt.Println(err)
  }

  return beer
}

func GetBeersByName(name string) []Beer {
  var beers []Beer

  err := postgres.Client.Model(&beers).
    Where("name LIKE ?", name).Select()

  if err != nil {
    fmt.Println(err)
  }

  return beers
}
