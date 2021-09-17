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

	if err = transaction.QueryRow(
		`INSERT INTO chat_rooms 
		DEFAULT VALUES 
		RETURNING chat_room_id, created_at, updated_at`,
	).
		Scan(
			&chatRoomModel.ChatRoomID,
			&chatRoomModel.CreatedAt,
			&chatRoomModel.UpdatedAt,
		); err != nil {
		transaction.Rollback()
		return err
	}

	for _, memberId := range memberIds {
		statement, err := transaction.Prepare(
			`INSERT INTO chat_room_members(chat_room_id, user_id) 
			VALUES($1, $2)`,
		)
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

func (r *Repository) GetChatRoomList(chatRoomPreview *dbModels.ChatRoomPreview, userId int64, limit int, offset int) error {
	statement, err := r.DatabaseClient.Prepare(
		`SELECT DISTINCT ON (CMembers.chat_room_id) CMembers.chat_room_id, Recipient.username as chat_room_name, CReads.total_unreads, CMsgs.message_type, CMsgs.message_content, Sender.username AS sender_name, CMsgs.created_at
		FROM chat_room_members AS CMembers
		JOIN chat_room_messages AS CMsgs ON CMsgs.chat_room_id = CMembers.chat_room_id
		JOIN (
			SELECT users.username AS username, chat_room_id
			FROM chat_room_members
			JOIN users ON users.user_id = chat_room_members.user_id AND users.user_id <> $1
		) AS Recipient ON Recipient.chat_room_id = CMembers.chat_room_id
		JOIN users AS Sender ON Sender.user_id = CMsgs.sender_id
		JOIN (
			SELECT COUNT(message_id) AS total_unreads, chat_room_id
			FROM chat_read_histories
			WHERE read_by <> $1
			GROUP BY chat_room_id
		) AS CReads ON CReads.chat_room_id = CMembers.chat_room_id
		WHERE CMembers.user_id = $1
		ORDER BY CMembers.chat_room_id, CMsgs.created_at DESC
		LIMIT $2 OFFSET $3`,
	)
	if err != nil {
		return err
	}

	if err = statement.QueryRow(userId, limit, offset).
		Scan(
			&chatRoomPreview.ChatRoomID,
			&chatRoomPreview.ChatRoomName,
			&chatRoomPreview.TotalUnreads,
			&chatRoomPreview.PreviewMessage.MessageType,
			&chatRoomPreview.PreviewMessage.MessageContent,
			&chatRoomPreview.PreviewMessage.SenderName,
			&chatRoomPreview.PreviewMessage.SentAt,
		); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetChatRoomMessages(chatRoomDetail *dbModels.ChatRoomDetail, userId int64, limit int, offset int) error {
	statement, err := r.DatabaseClient.Prepare(
		`SELECT *
		FROM chat_room_messages
		WHERE chat_room_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`,
	)
	if err != nil {
		return err
	}

	rows, err := statement.Query(chatRoomDetail.ChatRoomID, limit, offset)
	if err != nil {
		return err
	}

	if err = rows.Scan(&chatRoomDetail.Messages); err != nil {
		return err
	}

	statement, err = r.DatabaseClient.Prepare(
		`SELECT U.username
		FROM users AS U
		JOIN chat_room_members AS CMembers ON CMembers.user_id = U.user_id
		WHERE CMembers.chat_room_id = $1 AND CMembers.user_id <> $2`,
	)
	if err != nil {
		return err
	}

	if err = statement.QueryRow(chatRoomDetail.ChatRoomID, userId).Scan(&chatRoomDetail.ChatRoomName); err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateChatMessage(chatMessage *dbModels.ChatRoomMessage) error {
	transaction, err := r.DatabaseClient.Begin()
	if err != nil {
		transaction.Rollback()
		return err
	}

	statement, err := transaction.Prepare(
		`INSERT INTO chat_room_messages(chat_room_id, sender_id, message_type, message_content) 
		VALUES($1, $2, $3, $4) 
		RETURNING chat_room_message_id`,
	)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if err := statement.QueryRow(
		chatMessage.ChatRoomID,
		chatMessage.SenderID,
		chatMessage.MessageType,
		chatMessage.MessageContent,
	).Scan(&chatMessage.ChatRoomMessageID); err != nil {
		transaction.Rollback()
		return err
	}

	statement, err = transaction.Prepare(
		`INSERT INTO chat_read_histories(chat_room_id, message_id, read_by) 
		VALUES($1, $2, $3)`,
	)
	if err != nil {
		transaction.Rollback()
		return err
	}

	_, err = statement.Exec(chatMessage.ChatRoomID, chatMessage.ChatRoomMessageID, chatMessage.SenderID)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateReadHistory(readHistory *dbModels.ChatReadHistory) error {
	statement, err := r.DatabaseClient.Prepare(
		`INSERT INTO chat_read_histories(chat_room_id, message_id, read_by)
		 VALUES($1, $2, $3)`,
	)
	if err != nil {
		return err
	}

	_, err = statement.Exec(readHistory.ChatRoomID, readHistory.MessageID, readHistory.ReadBy)
	if err != nil {
		return err
	}

	return nil
}
