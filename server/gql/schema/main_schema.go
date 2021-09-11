package gqlSchema

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	gqlTypes "github.com/wildanpurnomo/go-persistent-chat/server/gql/types"
)

func CreateGQLRequestHandler() (gin.HandlerFunc, error) {
	schemaConfig := graphql.SchemaConfig{
		Query:    gqlTypes.QueryType,
		Mutation: gqlTypes.MutationType,
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	handler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return gin.HandlerFunc(func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}), nil
}
