package gql

import (
  "github.com/graphql-go/graphql"
)

type Root struct {
  Query *graphql.Object
}

func NewRoot() *Root {
  root := Root{
    Query: graphql.NewObject(
      graphql.ObjectConfig{
        Name: "Query",
        Fields: graphql.Fields{
          "beers": &graphql.Field{
            Type: graphql.NewList(Beer),
            Args: graphql.FieldConfigArgument{
              "name": &graphql.ArgumentConfig{
                Type: graphql.String,
              },
            },
            Resolve: BeerResolver,
          },
        },
      },
    ),
  }

  return &root
}
