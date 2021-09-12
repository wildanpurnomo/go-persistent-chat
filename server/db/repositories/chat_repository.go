package dbRepositories

import (
	dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"
)

func (r *Repository) CreateNewChatRoom(chatRoomModel *dbModels.ChatRoom, memberIds ...int64) error {
	transaction, err := r.DatabaseClient.Begin()
	if err != nil {
		transaction.Rollback()
		return err
	}

	if err = transaction.QueryRow("INSERT INTO chat_rooms DEFAULT VALUES RETURNING chat_room_id, created_at, updated_at").
		Scan(
			&chatRoomModel.ChatRoomID,
			&chatRoomModel.CreatedAt,
			&chatRoomModel.UpdatedAt,
		); err != nil {
		transaction.Rollback()
		return err
	}

	for _, memberId := range memberIds {
		statement, err := transaction.Prepare("INSERT INTO chat_room_members(chat_room_id, user_id) VALUES($1, $2)")
		if err != nil {
			transaction.Rollback()
			return err
		}

		if _, err = statement.Exec(chatRoomModel.ChatRoomID, memberId); err != nil {
			transaction.Rollback()
			return err
		}
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetChatRoomByMemberIds(chatRoomModel *dbModels.ChatRoom, memberIds ...int64) error {
	statement, err := r.DatabaseClient.Prepare(
		`SELECT C0.chat_room_id, C0.created_at, C0.updated_at 
		FROM chat_room_members AS C0
		JOIN chat_room_members AS C1 ON C1.chat_room_id = C0.chat_room_id AND C1.user_id = $2
		WHERE C0.deleted_at IS NULL AND C0.user_id = $1`,
	)
	if err != nil {
		return err
	}

	if err = statement.QueryRow(memberIds[0], memberIds[1]).
		Scan(
			&chatRoomModel.ChatRoomID,
			&chatRoomModel.CreatedAt,
			&chatRoomModel.UpdatedAt,
		); err != nil {
		return err
	}

	return nil
}
