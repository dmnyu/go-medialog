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
	ID       int    `json:"id"`
	AspaceID int    `json:"aspace_id"`
	Name     string `json:"name"`
}

type CreateRepository struct {
	AspaceID int    `json:"aspace_id"`
	Name     string `json:"name"`
}

type Resource struct {
	gorm.Model
	ID           int `json:"id"`
	ResourceID   int `json:"resource_id"`
	RepositoryID int `json:"repository_id"`
}

type CreateResource struct {
	ResourceID   int `json:"resource_id"`
	RepositoryID int `json:"repository_id"`
}

type Accession struct {
	gorm.Model
	ID           int `json:"id"`
	RepositoryID int `json:"repository_id"`
	AccessionID  int `json:"accession_id"`
}

type CreateAccession struct {
	RepositoryID int `json:"repository_id"`
	AccessionID  int `json:"accession_id"`
}

type Entry struct {
	gorm.Model
	ID            uuid.UUID `json:"id"`
	RepositoryID  int       `json:"repository_id"`
	ResourceID    int       `json:"resource_id"`
	AccessionID   int       `json:"accession_id"`
	MediaID       int       `json:"media_id"`
	MediaType     string    `json:"media_type"`
	BoxNumber     string    `json:"box_number"`
	LabelText     string    `json:"label_text"`
	OriginalID    string    `json:"original_id"`
	RefID         string    `json:"ref_id"`
	MediaNote     string    `json:"media_note"`
	StockUnit     string    `json:"stock_unit"`
	StockSize     int       `json:"stock_size"`
	IsRefreshed   bool      `json:"is_refreshed"`
	IsTransferred bool      `json:"is_transferred"`
}
