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
	ID           int    `json:"id"`
	AspaceID     int    `json:"aspace_id"`
	RepositoryID int    `json:"repository_id"`
	Name         string `json:"name"`
}

type CreateResource struct {
	AspaceID     int    `json:"aspace_id"`
	RepositoryID int    `json:"repository_id"`
	Name         string `json:"name"`
}

type Accession struct {
	ID				int `json:"id"`
	RepositoryID	int	`json:repository_id`
	AccessionID		int `json:"accession_id"`
}
