package index

type ESShard struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type ESTotal struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type ESMedia struct {
	ID           int    `json:"id"`
	CreatedAT    string `json:"CreatedAt"`
	UpdatedAt    string `json:"UpdatedAt"`
	DeletedAt    string `json:"DeletedAt"`
	ObjectID     int    `json:"object_id"`
	ModelID      int    `json:"model_id"`
	MediaID      int    `json:"media_id"`
	Subtype      string `json:"subtype"`
	HumanSize    string `json:"human_size"`
	RepositoryID int    `json:"repository_id"`
	ResourceID   int    `json:"resource_id"`
	AccessionID  int    `json:"accession_id"`
}

type ESHits struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source ESMedia `json:"_source"`
}

type ESHitsContainer struct {
	Total    ESTotal `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []ESHits
}

type ESResponse struct {
	Took     int             `json:"took"`
	TimedOut bool            `json:"timed_out"`
	Shards   ESShard         `json:"_shards"`
	Hits     ESHitsContainer `json:"hits"`
}
