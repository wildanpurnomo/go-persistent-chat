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
				"me": &graphql.Field{
					Type:        UserType,
					Description: "Get requesting user data",
					Resolve:     gqlResolvers.AuthenticateResolver,
				},
			},
		},
	)
)
