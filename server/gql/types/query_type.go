package gqlTypes

import (
	"github.com/graphql-go/graphql"
)

var (
	QueryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "QueryType",
			Fields: graphql.Fields{
				"echo": &graphql.Field{
					Type:        graphql.String,
					Description: "Echo",
					Args: graphql.FieldConfigArgument{
						"msg": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Args["msg"].(string), nil
					},
				},
			},
		},
	)
)
