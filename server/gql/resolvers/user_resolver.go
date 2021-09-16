package gqlResolvers

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/wildanpurnomo/go-persistent-chat/server/controllers"
	dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"
	authLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/auth"
	logsLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/logs"
)

var (
	AuthenticateResolver = func(params graphql.ResolveParams) (interface{}, error) {
		isValidRequest, userId := authLibs.CheckWhetherRequestIsUsingValidToken(params.Context)
		if !isValidRequest {
			return nil, errors.New("Invalid HTTP Request")
		}

		var userModel dbModels.User
		userModel.UserID = userId

		if err := controllers.Authenticate(&userModel); err != nil {
			return nil, err
		}

		return userModel, nil
	}
	RegisterResolver = func(params graphql.ResolveParams) (interface{}, error) {
		username := params.Args["username"].(string)
		password := params.Args["password"].(string)

		userModel := dbModels.User{Username: username, Password: password}

		if err := controllers.Register(&userModel); err != nil {
			logsLibs.LogIfDebug(err.Error())
			return nil, err
		}

		authLibs.SetJwtCookie(params.Context, userModel.UserID)

		return userModel, nil
	}
	LoginResolver = func(params graphql.ResolveParams) (interface{}, error) {
		username := params.Args["username"].(string)
		password := params.Args["password"].(string)

		userModel, err := controllers.Login(username, password)
		if err != nil {
			logsLibs.LogIfDebug(err.Error())
			return nil, err
		}

		authLibs.SetJwtCookie(params.Context, userModel.UserID)

		return userModel, nil
	}
	LogoutResolver = func(params graphql.ResolveParams) (interface{}, error) {
		authLibs.ClearJwtCookie(params.Context)
		return true, nil
	}
)
