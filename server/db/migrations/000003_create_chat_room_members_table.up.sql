CREATE TABLE IF NOT EXISTS chat_room_members(
    chat_room_member_id SERIAL PRIMARY KEY,
    chat_room_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

ALTER TABLE chat_room_members ADD FOREIGN KEY (chat_room_id) REFERENCES chat_rooms(chat_room_id);

ALTER TABLE chat_room_members ADD FOREIGN KEY (user_id) REFERENCES users(user_id);

CREATE INDEX idx_chat_room_members ON chat_room_members(chat_room_id);