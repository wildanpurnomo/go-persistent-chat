CREATE TABLE IF NOT EXISTS chat_read_histories(
    chat_read_history_id SERIAL PRIMARY KEY,
    chat_room_id INTEGER NOT NULL,
    message_id INTEGER NOT NULL,
    read_by INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

ALTER TABLE chat_read_histories ADD FOREIGN KEY (chat_room_id) REFERENCES chat_rooms(chat_room_id);

ALTER TABLE chat_read_histories ADD FOREIGN KEY (message_id) REFERENCES chat_room_messages(chat_room_message_id);

ALTER TABLE chat_read_histories ADD FOREIGN KEY (read_by) REFERENCES users(user_id);

CREATE INDEX idx_chat_read_histories ON chat_read_histories(chat_room_id, read_by);