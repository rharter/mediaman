package tmdb

type Episode struct {
	Id             int64   `json:"id"`
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int64   `json:"episode_number"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int64   `json:"season_number"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int64   `json:"vote_count"`
}
