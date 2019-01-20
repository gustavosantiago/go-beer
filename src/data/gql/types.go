package gql

import "github.com/graphql-go/graphql"

var Beer = graphql.NewObject(
  graphql.ObjectConfig{
    Name: "Beer",
    Fields: graphql.Fields{
      "id": &graphql.Field{
        Type: graphql.Int,
      },
      "name": &graphql.Field{
        Type: graphql.String,
      },
      "ibu": &graphql.Field{
        Type: graphql.Float,
      },
      "abv": &graphql.Field{
        Type: graphql.Float,
      },
      "description": &graphql.Field{
        Type: graphql.String,
      },
      "updated_at": &graphql.Field{
        Type: graphql.String,
      },
    },
  },
)
