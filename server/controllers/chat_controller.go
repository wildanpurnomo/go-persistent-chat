package controllers

import (
	"database/sql"

	"github.com/wildanpurnomo/go-persistent-chat/server/db"
	dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"
)

func CreateChatRoom(memberIds ...int64) (*dbModels.ChatRoom, error) {
	chatRoomModel := dbModels.ChatRoom{}

	isChatRoomExist, err := IsChatRoomForMemberIdsHasBeenCreated(&chatRoomModel, memberIds...)
	if err != nil {
		return nil, err
	}

	if isChatRoomExist {
		return &chatRoomModel, nil
	}

	if err := db.AppRepository.CreateNewChatRoom(&chatRoomModel, memberIds...); err != nil {
		return nil, err
	}

	return &chatRoomModel, nil
}

func IsChatRoomForMemberIdsHasBeenCreated(chatRoomModel *dbModels.ChatRoom, memberIds ...int64) (bool, error) {
	err := db.AppRepository.GetChatRoomByMemberIds(chatRoomModel, memberIds...)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	} else if err == sql.ErrNoRows {
		return false, nil
	} else {
		return true, nil
	}
}
