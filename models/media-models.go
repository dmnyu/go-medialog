package models

import "gorm.io/gorm"

type MediaModel int

const (
	OpticalDisc MediaModel = iota
	HardDiskDrive
)

type SystemIdentifiers struct {
	RepositoryID int
	ResourceID   int
	AccessionID  int
}

type Media interface {
	getIdentifiers() SystemIdentifiers
	getSizeInBytes() int64
	getHumanSize() string
}

type MediaOptical struct {
	gorm.Model
	RepositoryID int
	ResourceID   int
	AccessionID  int
	Manufacturer string
	StockSize    int
	StockUnit    string
	SizeInBytes  int64
	OpticalType  string
}

func (o MediaOptical) getIdentifiers() SystemIdentifiers {
	return SystemIdentifiers{
		RepositoryID: o.RepositoryID,
		ResourceID:   o.ResourceID,
		AccessionID:  o.AccessionID,
	}
}

func (o MediaOptical) getSizeInBytes() int64 {
	return o.SizeInBytes
}

type MediaHardDrive struct {
	gorm.Model
	RepositoryID int
	ResourceID   int
	AccessionID  int
	Manufacturer string
	StockSize    int
	StockUnit    string
	SizeInBytes  int64
	SerialNumber string
}

func (h MediaHardDrive) getIdentifiers() SystemIdentifiers {
	return SystemIdentifiers{
		RepositoryID: h.RepositoryID,
		ResourceID:   h.ResourceID,
		AccessionID:  h.AccessionID,
	}
}

func (h MediaHardDrive) getSizeInBytes() int64 {
	return h.SizeInBytes
}
