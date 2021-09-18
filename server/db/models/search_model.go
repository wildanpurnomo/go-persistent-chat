package dbModels

type UserSearch struct {
	SearchQuery  string             `json:"search_query"`
	SearchResult []UserSearchResult `json:"search_result"`
	HasMore      bool               `json:"has_more"`
	Limit        int
	Offset       int
}

type UserSearchResult struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	TotalRows int    `json:"total_rows"`
}

func (s *UserSearch) ResultNotEmpty() bool {
	return len(s.SearchResult) > 0
}

func (s *UserSearch) DetermineHasMore() {
	totalRows := s.SearchResult[0].TotalRows

	if s.Limit > totalRows {
		s.HasMore = false
	} else if totalRows%s.Limit == 0 {
		s.HasMore = s.Offset < totalRows/s.Limit-1
	} else {
		s.HasMore = s.Offset < (totalRows+s.Limit-1)/s.Limit
	}
}

func (s *UserSearch) GetWildcardQuery() string {
	return "%" + s.SearchQuery + "%"
}
