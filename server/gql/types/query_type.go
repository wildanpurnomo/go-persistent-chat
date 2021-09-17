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
				"search_users": &graphql.Field{
					Type:        SearchUsersType,
					Description: "Search users by their username",
					Args: graphql.FieldConfigArgument{
						"search_query": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"offset": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: gqlResolvers.SearchUsersResolver,
				},
			},
		},
	)
)
