package gqlResolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/wildanpurnomo/go-persistent-chat/server/controllers"
	dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"
	authLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/auth"
	logsLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/logs"
)

var (
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
)
