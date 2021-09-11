package dbRepositories

import (
	dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"
)

func (r *Repository) CreateNewUser(userModel *dbModels.User) error {
	statement, err := r.DatabaseClient.Prepare("INSERT INTO users(username, password) VALUES($1, $2) RETURNING user_id, created_at, updated_at")
	if err != nil {	
		return err
	}

	if err := statement.QueryRow(
		userModel.Username,
		userModel.Password).
		Scan(
			&userModel.UserID,
			&userModel.CreatedAt,
			&userModel.UpdatedAt,
		); err != nil {
		return err
	}

	return nil
}

func (r *Repository) SearchUserByUsername(userModel *dbModels.User) error {
	statement, err := r.DatabaseClient.Prepare("SELECT user_id, username, password, created_at, updated_at FROM users WHERE deleted_at IS NULL AND username=$1")
	if err != nil {
		return err
	}

	if err = statement.QueryRow(
		userModel.Username).
		Scan(
			&userModel.UserID,
			&userModel.Username,
			&userModel.Password,
			&userModel.CreatedAt,
			&userModel.UpdatedAt,
		); err != nil {
		return err
	}

	return nil
}
