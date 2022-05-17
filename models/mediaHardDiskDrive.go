package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

var HardDriveSubtypes = []string{"5.25 inch magnetic", "3.5 inch magnetic", "3.5 inch SSD"}

type MediaHardDrive struct {
	gorm.Model
	RepositoryID int    `json:"repository_id" form:"repository_id"`
	ResourceID   int    `json:"resource_id" form:"resource_id"`
	AccessionID  int    `json:"accession_id" form:"accession_id"`
	Internal     bool   `json:"internal" form:"internal"`
	Manufacturer string `json:"manufacturer" form:"manufacturer"`
	MediaID      int    `json:"media_id" form:"media_id"`
	MediaNote    string `json:"media_note" form:"media_note"`
	StockSize    int    `json:"stock_size" form:"stock_size"`
	StockUnit    string `json:"stock_unit" form:"stock_unit"`
	SizeInBytes  int64  `json:"size_in_bytes" form:"size_in_bytes"`
	SerialNumber string `json:"serial_number" form:"serial_number"`
	Subtype      string `json:"subtype" form:"subtype"`
}

func (h MediaHardDrive) GetMediaEntry() MediaEntry {
	j, _ := json.Marshal(h)

	return MediaEntry{
		ModelID:      HardDiskDrive,
		MediaID:      h.MediaID,
		DatabaseID:   h.ID,
		Subtype:      h.Subtype,
		HumanSize:    h.GetHumanSize(),
		RepositoryID: h.RepositoryID,
		ResourceID:   h.ResourceID,
		AccessionID:  h.AccessionID,
		JSON:         string(j),
	}
}

func (h MediaHardDrive) GetIdentifiers() SystemIdentifiers {
	return SystemIdentifiers{
		RepositoryID: h.RepositoryID,
		ResourceID:   h.ResourceID,
		AccessionID:  h.AccessionID,
	}
}

func (h MediaHardDrive) GetSizeInBytes() int64 {
	return h.SizeInBytes
}

func (o MediaHardDrive) GetHumanSize() string {
	return fmt.Sprintf("%d %s", o.StockSize, o.StockUnit)
}
