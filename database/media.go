package database

import "gorm.io/gorm"

type MediaModel int

const (
	OpticalDisc MediaModel = iota
	HardDiskDrive
)

var SubTypes = map[int][]string{
	0: []string{"CD", "DVD"},
	1: []string{"3.5 in. Magnetic", "2.5 in. Magentic"},
}

type MediaObject struct {
	gorm.Model
	RepositoryID int    `json:"repository_id" form:"repository_id"`
	ResourceID   int    `json:"resource_id" form:"resource_id"`
	AccessionID  int    `json:"accession_id" form:"accession_id"`
	ModelID      int    `json:"model_id" form:"model_id"`
	ObjectID     int    `json:"object_id" form:"object_id"`
	Subtype      string `json:"subtype" form:"subtype"`
}

type MediaOpticalDisc struct {
	gorm.Model
	ModelID     MediaModel `json:"model_id" form:"model_id"`
	StockUnit   string     `json:"stock_unit" form:"stock_unit"`
	StockSize   int        `json:"stock_size" form:"stockSize"`
	SizeInBytes int64      `json:"size_in_bytes" form:size_in_bytes`
	Subtype     string     `json:"subtype"`
}

type MediaHardDiskDrive struct {
	gorm.Model
	ModelID   MediaModel
	StockUnit string
	StockSize int
	SubTypeID int
}
