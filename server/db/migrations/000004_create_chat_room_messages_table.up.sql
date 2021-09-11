CREATE TYPE CHAT_MESSAGE_TYPE AS ENUM('text', 'image', 'video', 'file');

CREATE TABLE IF NOT EXISTS chat_room_messages(
    chat_room_message_id SERIAL PRIMARY KEY,
    chat_room_id INTEGER NOT NULL,
    sender_id INTEGER NOT NULL,
    message_type CHAT_MESSAGE_TYPE NOT NULL,
    message_content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

ALTER TABLE chat_room_messages ADD FOREIGN KEY (chat_room_id) REFERENCES chat_rooms(chat_room_id);

ALTER TABLE chat_room_messages ADD FOREIGN KEY (sender_id) REFERENCES users(user_id);

CREATE INDEX idx_chat_room_messages ON chat_room_messages(chat_room_id);
