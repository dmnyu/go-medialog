package models

import (
	"github.com/google/uuid"
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
	ID       int
	AspaceID int    `json:"aspace_id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
}

type CreateRepository struct {
	AspaceID int    `json:"aspace_id" form:"aspace_id"`
	Slug    string `json:"slug" form:"slug"`
	Name     string `json:"name" form:"name"`
}

type Resource struct {
	gorm.Model
	ID                        int
	AspaceResourceID          int        `json:"resource_id"`
	RepositoryID              int        `json:"repository_id"`
	Repository                Repository `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AspaceResourceTitle       string     `json:"resource_title"`
	AspaceResourceIdentifiers string     `json:"resource_identifiers"`
}

type CreateResource struct {
	AspaceResourceID   int `json:"aspace_id"`
	AspaceRepositoryID int `json:"repository_id"`
}

type Accession struct {
	gorm.Model
	ID           int
	AspaceID     int        `json:"aspace_id"`
	RepositoryID int        `json:"repository_id"`
	Repository   Repository `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ResourceID   int        `json:"resource_id"`
	Resource     Resource   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title        string     `json:"accession_title"`
	Identifiers  string     `json:"identifier"`
	State        string     `json:"state"`
	CreatedBy    int        `json:"created_by"`
	UpdatedBy    int        `json:"updated_by"`
}

type CreateAccession struct {
	AspaceID           int `json:"aspace_id" form:"aspace_id"`
	AspaceRepositoryID int `json:"aspace_repository_id" form:"aspace_repository_id"`
}

type Entry struct {
	gorm.Model
	ID            uuid.UUID  `json:"id"`
	RepositoryID  int        `json:"repository_id"`
	Repository    Repository `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ResourceID    int        `json:"resource_id"`
	Resource      Resource   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AccessionID   int        `json:"accession_id"`
	Accession     Accession  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MediaID       int        `json:"media_id"`
	MediaType     string     `json:"media_type"`
	BoxNumber     string     `json:"box_number"`
	LabelText     string     `json:"label_text"`
	OriginalID    string     `json:"original_id"`
	RefID         string     `json:"ref_id"`
	MediaNote     string     `json:"media_note"`
	StockUnit     string     `json:"stock_unit"`
	StockSize     int        `json:"stock_size"`
	IsRefreshed   bool       `json:"is_refreshed"`
	IsTransferred bool       `json:"is_transferred"`
	MediaModel		int		`json:"media_model"`
	MediaModelID	int	`json:"media_model_id"`
}

type MediaModel int

const (
	OpticalDisc MediaModel = iota
	HardDiskDrive
)
