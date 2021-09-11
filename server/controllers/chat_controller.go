package controllers

import (
	"github.com/wildanpurnomo/go-persistent-chat/server/db"
	dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"
)

func CreateChatRoom(memberIds ...int64) (*dbModels.ChatRoom, error) {
	chatRoomModel := dbModels.ChatRoom{}

	if err := db.AppRepository.CreateNewChatRoom(&chatRoomModel, memberIds...); err != nil {
		return nil, err
	}

	return &chatRoomModel, nil
}
