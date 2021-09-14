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

type ChatRoomPreview struct {
	ChatRoomID     int64                  `json:"chat_room_id"`
	ChatRoomName   string                 `json:"chat_room_name"`
	TotalUnreads   int                    `json:"total_unreads"`
	PreviewMessage ChatRoomPreviewMessage `json:"chat_room_preview_message"`
}

type ChatRoomPreviewMessage struct {
	MessageType    string `json:"message_type"`
	MessageContent string `json:"message_content"`
	SenderName     string `json:"sender_name"`
	SentAt         string `json:"sent_at"`
}

type ChatRoomDetail struct {
	ChatRoomID   int64             `json:"chat_room_id"`
	ChatRoomName string            `json:"chat_room_name"`
	Messages     []ChatRoomMessage `json:"chat_messages"`
}
