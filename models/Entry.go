package models

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	ID           int
	RepositoryId int
	ResourceId   int
	AccessionId  int
	MediaType    string
	BoxNumber    string
	Refreshed    bool
}
