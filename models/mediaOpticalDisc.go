package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

var OpticalSubtypes = []string{"CD", "CD-R", "CD-RW", "DVD", "DVD-R", "DVD-RW"}

var OpticalContentTypes = []string{"Audio", "Data", "Video", "Unknown"}

type MediaOpticalDisc struct {
	gorm.Model
	ModelID         MediaModel `json:"model_id" form:"model_id"`
	MediaID         int        `json:"media_id" form:"media_id"`
	RepositoryID    int        `json:"repository_id" form:"repository_id"`
	ResourceID      int        `json:"resource_id" form:"resource_id"`
	AccessionID     int        `json:"accession_id" form:"accession_id"`
	StockUnit       string     `json:"stock_unit" form:"stock_unit"`
	StockSize       int        `json:"stock_size" form:"stock_size"`
	SizeInBytes     int64      `json:"size_in_bytes" form:"size_in_bytes"`
	Subtype         string     `json:"subtype" form:"subtype"`
	Manufacturer    string     `json:"manufacturer" form:"manufacturer"`
	MediaNote       string     `json:"media_note" form:"media_note""`
	Diameter        float32    `json:"diameter" form:"diameter"`
	DispositionNote string     `json:"disposition_note" form:"disposition_note"`
	LabelText       string     `json:"label_text" form:"label_text"`
	ContentType     string     `json:"content_type" form:"content_type"`
	OriginalID      string     `json:"original_id" form:"original_id"`
	RefID           string     `json:"ref_id" form:"ref_id"`
	BoxNumber       string     `json:"box_number" form:"box_number"`
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

func (o MediaOpticalDisc) GetSizeInBytes() int64 {
	return int64(o.SizeInBytes)
}

func (o MediaOpticalDisc) GetHumanSize() string {
	return fmt.Sprintf("%d %s", o.StockSize, o.StockUnit)
}
