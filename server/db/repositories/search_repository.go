package dbRepositories

import dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"

func (r *Repository) SearchUsersByUsername(searchModel *dbModels.UserSearch) error {
	statement, err := r.DatabaseClient.Prepare(
		`SELECT user_id, username, COUNT(*) OVER() as total_rows
		FROM users
		WHERE deleted_at IS NULL AND username LIKE $1
		LIMIT $2 OFFSET $3`,
	)
	if err != nil {
		return err
	}

	rows, err := statement.Query(searchModel.GetWildcardQuery(), searchModel.Limit, searchModel.Offset)
	if err != nil {
		return err
	}

	for rows.Next() {
		var searchResult dbModels.UserSearchResult
		if err = rows.Scan(&searchResult.UserID, &searchResult.Username, &searchResult.TotalRows); err != nil {
			return err
		}

		searchModel.SearchResult = append(searchModel.SearchResult, searchResult)
	}

	return nil
}
