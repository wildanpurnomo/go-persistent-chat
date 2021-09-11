package gqlTypes

import (
	"github.com/graphql-go/graphql"
	gqlResolvers "github.com/wildanpurnomo/go-persistent-chat/server/gql/resolvers"
)

var (
	QueryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "QueryType",
			Fields: graphql.Fields{
				"login": &graphql.Field{
					Type:        UserType,
					Description: "Login as a user",
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: gqlResolvers.LoginResolver,
				},
			},
		},
	)
)
