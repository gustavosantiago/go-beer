package gql

import (
  beer "data/models"

  "github.com/graphql-go/graphql"
)

func BeerResolver(p graphql.ResolveParams) (interface{}, error) {

  name, ok := p.Args["name"].(string)

  if ok {
    result := beer.GetBeersByName(name)

    return result, nil
  } else {
    result := beer.GetBeers()

    return result, nil
  }

}
