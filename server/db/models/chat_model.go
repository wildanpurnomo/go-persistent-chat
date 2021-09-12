package dbModels

type ChatRoom struct {
	ChatRoomID int64  `json:"chat_room_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ChatRoomMember struct {
	ChatRoomMemberID int64  `json:"chat_room_member_id"`
	ChatRoomID       int64  `json:"chat_room_id"`
	UserID           int64  `json:"user_id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type ChatRoomMessage struct {
	ChatRoomMessageID int64  `json:"chat_room_message_id"`
	ChatRoomID        int64  `json:"chat_room_id"`
	SenderID          int64  `json:"sender_id"`
	MessageType       string `json:"message_type"`
	MessageContent    string `json:"message_content"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type ChatReadHistory struct {
	ChatReadHistoryID int64  `json:"chat_read_history_id"`
	ChatRoomID        int64  `json:"chat_room_id"`
	MessageID         int64  `json:"message_id"`
	ReadBy            int64  `json:"read_by"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}
