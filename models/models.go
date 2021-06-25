package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      int    `json:"id"`
	Email   string `json:"email"`
	PassMD5 string `json:"pass_md5"`
	Salt    string `json:"salt"`
	IsAdmin bool   `json:"is_admin"`
}

type Credentials struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
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
