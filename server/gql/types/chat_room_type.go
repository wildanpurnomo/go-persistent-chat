package gqlTypes

import "github.com/graphql-go/graphql"

var (
	ChatRoomType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "ChatRoomType",
			Fields: graphql.Fields{
				"chat_room_id": &graphql.Field{
					Type: graphql.Int,
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
				},
				"updated_at": &graphql.Field{
					Type: graphql.String,
				},
				"deleted_at": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
)
