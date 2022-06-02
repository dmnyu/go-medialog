package models

type MediaEntry struct {
	ModelID      MediaModel `json:"model_id" form:"model_id"`
	MediaID      int        `json:"media_id" form:"media_id"`
	DatabaseID   uint       `json:"database_id"`
	Subtype      string     `json:"subtype" form:"subtype"`
	HumanSize    string     `json:"human_size" form:"human_size"`
	RepositoryID int        `json:"repository_id" form:"repository_id"`
	ResourceID   int        `json:"resource_id" form:"resource_id"`
	AccessionID  int        `json:"accession_id" form:"accession_id"`
	JSON         string     `json:"json"`
}

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

type ESCreateResponse struct {
	Index       string  `json:"_index"`
	Type        string  `json:"_type"`
	ID          string  `json:"_id"`
	Version     int     `json:"_version"`
	Result      string  `json:"result"`
	Shards      ESShard `json:"_shards"`
	SeqNum      int     `json:"_seq_num"`
	PrimaryTerm int     `json:"_primary_term"`
}

type ESHit struct {
	Index  string     `json:"_index"`
	Type   string     `json:"_type"`
	ID     string     `json:"_id"`
	Score  float64    `json:"_score"`
	Source MediaEntry `json:"_source"`
}

type ESHitsContainer struct {
	Total    ESTotal `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []ESHit
}

type ESResponse struct {
	Took     int             `json:"took"`
	TimedOut bool            `json:"timed_out"`
	Shards   ESShard         `json:"_shards"`
	Hits     ESHitsContainer `json:"hits"`
}

type IndexQuery struct {
	Query string `json:"query" form:"query"`
}
