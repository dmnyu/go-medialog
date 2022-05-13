package models

import (
	"gorm.io/gorm"
)

type Credentials struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type User struct {
	gorm.Model
	FirstName  string `json:"first_name" form:"first_name"`
	LastName   string `json:"last_name" form:"last_name"`
	Email      string `json:"email" form:"email" gorm:"uniqueIndex"`
	PassSHA512 string `json:"pass_md5"`
	Salt       string `json:"salt"`
	IsAdmin    bool   `json:"is_admin" form:"is_admin"`
}

type CreateUser struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Email     string `json:"email" form:"email"`
	Password1 string `json:"password_1" form:"password_1"`
	Password2 string `json:"password_2" form:"password_2"`
	IsAdmin   bool   `json:"is_admin" form:"is_admin"`
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
