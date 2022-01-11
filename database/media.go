package database

import (
	"fmt"
	"gorm.io/gorm"
)

type MediaModel int

const (
	OpticalDisc MediaModel = iota
	HardDiskDrive
)

var SubTypes = map[int][]string{
	0: {"CD", "DVD"},
	1: {"3.5 in. Magnetic", "2.5 in. Magnetic", "2.5 in. SSD"},
}

type MediaCore struct {
	ID           uint
	ModelID      MediaModel
	Subtype      string
	HumanSize    string
	RepositoryID int
	ResourceID   int
	AccessionID  int
}

type Media interface {
	getMediaCore() MediaCore
}

type MediaEntry struct {
	gorm.Model
	ModelID      MediaModel
	MediaID      int
	RepositoryID int
	ResourceID   int
	AccessionID  int
}

type MediaOpticalDisc struct {
	gorm.Model
	ModelID      MediaModel `json:"model_id" form:"model_id"`
	RepositoryID int
	ResourceID   int
	AccessionID  int
	StockUnit    string `json:"stock_unit" form:"stock_unit"`
	StockSize    int    `json:"stock_size" form:"stockSize"`
	SizeInBytes  int64  `json:"size_in_bytes" form:"size_in_bytes"`
	Subtype      string `json:"subtype"`
	MediaNote    string `json:"media_note"`
}

func (o MediaOpticalDisc) getMediaCore() MediaCore {
	return MediaCore{
		ID:           o.ID,
		ModelID:      o.ModelID,
		Subtype:      o.Subtype,
		HumanSize:    fmt.Sprintf("%d %s", o.StockSize, o.StockUnit),
		RepositoryID: o.RepositoryID,
		ResourceID:   o.ResourceID,
		AccessionID:  o.AccessionID,
	}
}

type MediaHardDiskDrive struct {
	gorm.Model
	ModelID      MediaModel
	RepositoryID int
	ResourceID   int
	AccessionID  int
	StockUnit    string
	StockSize    int
	SizeInBytes  int64 `json:"size_in_bytes" form:"size_in_bytes"`
	Subtype      string
	MediaNote    string
}

func (hd MediaHardDiskDrive) getMediaCore() MediaCore {
	return MediaCore{
		ID:           hd.ID,
		ModelID:      hd.ModelID,
		Subtype:      hd.Subtype,
		HumanSize:    fmt.Sprintf("%d %s", hd.StockSize, hd.StockUnit),
		RepositoryID: hd.RepositoryID,
		ResourceID:   hd.ResourceID,
		AccessionID:  hd.AccessionID,
	}
}
