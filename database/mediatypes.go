package database

import "gorm.io/gorm"

type MediaModel int

const (
	OpticalDisc MediaModel = iota
	HardDiskDrive
)

var subTypes = map[int][]string{
	0: []string{"CD", "DVD"},
	1: []string{"3.5 in. Magnetic", "2.5 in. Magentic"},
}

type MediaObject struct {
	gorm.Model
	RepositoryID int
	ResourceID   int
	AccessionID  int
	ModelID      int
	ObjectID     int
}

type MediaOpticalDisc struct {
	gorm.Model
	ModelID   MediaModel
	StockUnit string
	StockSize int
	SubType   string
}

type MediaHardDiskDrive struct {
	gorm.Model
	ModelID   MediaModel
	StockUnit string
	StockSize int
	SubTypeID int
}
