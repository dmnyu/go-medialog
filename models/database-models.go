package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type Credentials struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type User struct {
	gorm.Model
	ID      int    `json:"id"`
	Email   string `json:"email"`
	PassMD5 string `json:"pass_md5"`
	Salt    string `json:"salt"`
	IsAdmin bool   `json:"is_admin"`
}

type Repository struct {
	gorm.Model
	AspaceID int    `json:"aspace_id" form:"aspace_id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
}

type Resource struct {
	gorm.Model
	AspaceID     int    `json:"aspace_resource_id"`
	RepositoryID int    `json:"repository_id"`
	Title        string `json:"resource_title"`
	Identifiers  string `json:"resource_identifiers"`
}

type Accession struct {
	gorm.Model
	AspaceID     int    `json:"aspace_id"`
	RepositoryID int    `json:"repository_id"`
	ResourceID   int    `json:"resource_id"`
	Title        string `json:"title"`
	Identifiers  string `json:"identifier"`
	State        string `json:"state"`
}

type CreateAspaceObject struct {
	RepositoryID int `json:"repository_id" form:"repository_id"`
	ResourceID   int `json:"resource_id" form:"resource_id"`
	AccessionID  int `json:"accession_id" form:"accession_id"`
}

var SubTypes = map[int][]string{
	0: {"CD", "DVD"},
	1: {"3.5 in. Magnetic", "2.5 in. Magnetic", "2.5 in. SSD"},
}

type MediaOpticalDisc struct {
	gorm.Model
	ModelID      MediaModel `json:"model_id" form:"model_id"`
	MediaID      int        `json:"media_id" form:"media_id"`
	RepositoryID int        `json:"repository_id" form:"repository_id"`
	ResourceID   int        `json:"resource_id" form:"resource_id"`
	AccessionID  int        `json:"accession_id" form:"accession_id"`
	StockUnit    string     `json:"stock_unit" form:"stock_unit"`
	StockSize    int        `json:"stock_size" form:"stock_size"`
	SizeInBytes  int64      `json:"size_in_bytes" form:"size_in_bytes"`
	Subtype      string     `json:"subtype" form:"subtype"`
	MediaNote    string     `json:"media_note"`
}

func (o MediaOpticalDisc) GetMediaEntry() MediaEntry {
	j, _ := json.Marshal(o)

	return MediaEntry{
		ModelID:      0,
		MediaID:      o.MediaID,
		ObjectID:     int(o.ID),
		Subtype:      o.Subtype,
		HumanSize:    fmt.Sprintf("%d %s", o.StockSize, o.StockUnit),
		RepositoryID: o.RepositoryID,
		ResourceID:   o.ResourceID,
		AccessionID:  o.AccessionID,
		JSON:         string(j),
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
