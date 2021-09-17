package gqlTypes

import "github.com/graphql-go/graphql"

var (
	SearchUsersType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "SearchUsersType",
			Fields: graphql.Fields{
				"search_query": &graphql.Field{
					Type: graphql.String,
				},
				"has_more": &graphql.Field{
					Type: graphql.Boolean,
				},
				"search_result": &graphql.Field{
					Type: &graphql.List{
						OfType: UserType,
					},
				},
			},
		},
	)
)
