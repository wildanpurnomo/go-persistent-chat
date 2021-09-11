package gqlTypes

import "github.com/graphql-go/graphql"

var (
	UserType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UserType",
			Fields: graphql.Fields{
				"user_id": &graphql.Field{
					Type: graphql.Int,
				},
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
				},
				"updated_at": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
)