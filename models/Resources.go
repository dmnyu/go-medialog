package models

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	ID                 int
	AspaceResourceID   int
	AspaceRepositoryID int
	ResourceIds        string
	Title              string
}
