package controllers

import (
	"database/sql"

	"github.com/wildanpurnomo/go-persistent-chat/server/db"
	dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"
)

func CreateChatRoom(chatRoomModel *dbModels.ChatRoom, memberIds ...int64) error {
	isChatRoomExist, err := isChatRoomForMemberIdsHasBeenCreated(chatRoomModel, memberIds...)
	if err != nil {
		return err
	}

	if isChatRoomExist {
		return nil
	}

	if err := db.AppRepository.CreateNewChatRoom(chatRoomModel, memberIds...); err != nil {
		return err
	}

	return nil
}

func CreateChatMessage(chatMessage *dbModels.ChatRoomMessage) error {
	if err := db.AppRepository.CreateChatMessage(chatMessage); err != nil {
		return err
	}

	return nil
}

func CreateReadHistory(readHistory *dbModels.ChatReadHistory) error {
	if err := db.AppRepository.CreateReadHistory(readHistory); err != nil {
		return err
	}

	return nil
}

func GetChatRoomList(chatPreview *dbModels.ChatRoomPreview, userId int64, limit int, offset int) error {
	if err := db.AppRepository.GetChatRoomList(chatPreview, userId, limit, offset); err != nil {
		return err
	}

	return nil
}

func GetChatRoomMessages(chatRoomDetail *dbModels.ChatRoomDetail, userId int64, limit int, offset int) error {
	if err := db.AppRepository.GetChatRoomMessages(chatRoomDetail, userId, limit, offset); err != nil {
		return err
	}

	return nil
}

func isChatRoomForMemberIdsHasBeenCreated(chatRoomModel *dbModels.ChatRoom, memberIds ...int64) (bool, error) {
	err := db.AppRepository.GetChatRoomByMemberIds(chatRoomModel, memberIds...)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	} else if err == sql.ErrNoRows {
		return false, nil
	} else {
		return true, nil
	}
}
