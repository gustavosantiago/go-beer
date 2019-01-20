package postgres

import (
  "github.com/go-pg/pg"
)

func credentials() *pg.Options {
  return &pg.Options{
    Addr: "127.0.0.1:5432",
    User: "postgres",
    Database: "go_beer",}
}