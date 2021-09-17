package gqlResolvers

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/wildanpurnomo/go-persistent-chat/server/controllers"
	dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"
	authLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/auth"
)

var (
	SearchUsersResolver = func(params graphql.ResolveParams) (interface{}, error) {
		isValidRequest, _ := authLibs.CheckWhetherRequestIsUsingValidToken(params.Context)
		if !isValidRequest {
			return nil, errors.New("Invalid HTTP Request")
		}

		var searchModel dbModels.UserSearch
		searchModel.SearchQuery = params.Args["search_query"].(string)
		searchModel.Limit = params.Args["limit"].(int)
		searchModel.Offset = params.Args["offset"].(int)

		if err := controllers.SearchUsersByUsername(&searchModel); err != nil {
			return nil, err
		}

		if searchModel.ResultNotEmpty() {
			searchModel.DetermineHasMore()
		}

		return searchModel, nil
	}
)
