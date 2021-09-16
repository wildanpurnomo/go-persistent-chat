package gqlTypes

import (
	"github.com/graphql-go/graphql"
	gqlResolvers "github.com/wildanpurnomo/go-persistent-chat/server/gql/resolvers"
)

var (
	MutationType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "MutationType",
			Fields: graphql.Fields{
				"register": &graphql.Field{
					Type:        UserType,
					Description: "Register new user",
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: gqlResolvers.RegisterResolver,
				},
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
				"logout": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Log requesting user out",
					Resolve:     gqlResolvers.LogoutResolver,
				},
				"create_chat_room": &graphql.Field{
					Type:        ChatRoomType,
					Description: "Create new chat room",
					Args: graphql.FieldConfigArgument{
						"member_ids": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.Int),
						},
					},
					Resolve: gqlResolvers.CreateNewChatRoomResolver,
				},
			},
		},
	)
)
