package gqlResolvers

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/wildanpurnomo/go-persistent-chat/server/controllers"
	authLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/auth"
	logLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/logs"
)

var (
	CreateNewChatRoomResolver = func(params graphql.ResolveParams) (interface{}, error) {
		isValidRequest, userId := authLibs.CheckWhetherRequestIsUsingValidToken(params.Context)
		if !isValidRequest {
			return nil, errors.New("Invalid HTTP Request")
		}

		memberIds := []int64{userId}
		inputIds := params.Args["member_ids"].([]interface{})
		for _, inputId := range inputIds {
			memberIds = append(memberIds, int64(inputId.(int)))
		}

		chatRoomModel, err := controllers.CreateChatRoom(memberIds...)
		if err != nil {
			logLibs.LogIfDebug(err.Error())
			return nil, errors.New("Chat Room Creation Is Failed")
		}

		return chatRoomModel, nil
	}
)
