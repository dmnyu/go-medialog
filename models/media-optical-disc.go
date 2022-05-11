package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type MediaOpticalDisc struct {
	gorm.Model
	ModelID      MediaModel `json:"model_id" form:"model_id"`
	MediaID      int        `json:"media_id" form:"media_id"`
	RepositoryID int        `json:"repository_id" form:"repository_id"`
	ResourceID   int        `json:"resource_id" form:"resource_id"`
	AccessionID  int        `json:"accession_id" form:"accession_id"`
	StockUnit    string     `json:"stock_unit" form:"stock_unit"`
	StockSize    int        `json:"stock_size" form:"stock_size"`
	SizeInBytes  float64    `json:"size_in_bytes" form:"size_in_bytes"`
	Subtype      string     `json:"subtype" form:"subtype"`
	Manufacturer string     `json:"manufacturer" form:"manufacturer"`
	MediaNote    string     `json:"media_note"`
}

func (o MediaOpticalDisc) GetMediaEntry() MediaEntry {
	j, _ := json.Marshal(o)

	return MediaEntry{
		ModelID:      OpticalDisc,
		MediaID:      o.MediaID,
		DatabaseID:   o.ID,
		Subtype:      o.Subtype,
		HumanSize:    o.GetHumanSize(),
		RepositoryID: o.RepositoryID,
		ResourceID:   o.ResourceID,
		AccessionID:  o.AccessionID,
		JSON:         string(j),
	}
}
func (o MediaOpticalDisc) GetIdentifiers() SystemIdentifiers {
	return SystemIdentifiers{
		RepositoryID: o.RepositoryID,
		ResourceID:   o.ResourceID,
		AccessionID:  o.AccessionID,
	}
}

func (o MediaOpticalDisc) GetSizeInBytes() float64 {
	return o.SizeInBytes
}

func (o MediaOpticalDisc) GetHumanSize() string {
	return fmt.Sprintf("%d %s", o.StockSize, o.StockUnit)
}
