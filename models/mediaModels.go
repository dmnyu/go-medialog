package models

import "github.com/nyudlts/bytemath"

type MediaModel int

const (
	OpticalDisc MediaModel = iota
	HardDiskDrive
)

var MediaSubTypes = map[MediaModel][]string{}

var MediaNames = map[MediaModel]string{
	OpticalDisc:   "Optical Disc",
	HardDiskDrive: "Hard Disk Drive",
}

type MediaType struct {
	Name     string
	Subtypes []string
}

func GetMediaTypes() map[MediaModel]MediaType {
	mediaTypes := map[MediaModel]MediaType{}
	for k, v := range MediaSubTypes {
		mediaTypes[k] = MediaType{MediaNames[k], v}
	}
	return mediaTypes
}

var MediaUnit = []string{"B", "KB", "MB", "GB", "TB"}

var ByteMathSuffix = map[string]bytemath.Suffix{
	"B":  bytemath.B,
	"KB": bytemath.KB,
	"MB": bytemath.MB,
	"GB": bytemath.GB,
	"TB": bytemath.TB,
}

type SystemIdentifiers struct {
	RepositoryID int
	ResourceID   int
	AccessionID  int
}

type Media interface {
	GetMediaEntry() MediaEntry
	GetIdentifiers() SystemIdentifiers
	GetSizeInBytes() int64
	GetHumanSize() string
}
