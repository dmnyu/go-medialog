package database

import (
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
	AspaceID int    `json:"aspace_id"`
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

type MediaModel int

const (
	OpticalDisc MediaModel = iota
	HardDiskDrive
)
