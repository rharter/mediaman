package tmdb

import (
	"encoding/json"
)

type SearchResult struct {

	// Current Page
	Page int64 `json:"page"`

	TotalPages int64 `json:"total_pages"`

	TotalResults int64 `json:"total_results"`

	ResultsRaw json.RawMessage `json:"results"`
}
