package controllers

import (
	"github.com/wildanpurnomo/go-persistent-chat/server/db"
	dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"
)

func SearchUsersByUsername(searchModel *dbModels.UserSearch) error {
	if err := db.AppRepository.SearchUsersByUsername(searchModel); err != nil {
		return err
	}

	return nil
}
