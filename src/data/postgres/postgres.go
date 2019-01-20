package postgres

import(
  "github.com/go-pg/pg"
)

type DB pg.DB
type Options pg.Options

var Client = pg.Connect(credentials())