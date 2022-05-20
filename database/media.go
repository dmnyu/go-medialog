package database

import (
	"github.com/dmnyu/go-medialog/models"
)

func InsertOpticalDisc(disc *models.MediaOpticalDisc) error {
	if err := db.Create(disc).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOpticalDisc(discID uint) error {
	if err := db.Delete(&models.MediaOpticalDisc{}, discID).Error; err != nil {
		return err
	}
	return nil
}

func FindOpticalDisc(discID int) (models.MediaOpticalDisc, error) {
	o := models.MediaOpticalDisc{}
	if err := db.Find(&o, discID).Error; err != nil {
		return o, err
	}
	return o, nil
}

func FindOpticaDiscs() *[]models.MediaOpticalDisc {
	discs := []models.MediaOpticalDisc{}
	db.Find(&discs)
	return &discs
}

func InsertHardDiskDrive(hd *models.MediaHardDrive) error {
	if err := db.Create(hd).Error; err != nil {
		return err
	}
	return nil
}

func DeleteHardDiskDrive(hddID uint) error {
	if err := db.Delete(&models.MediaOpticalDisc{}, hddID).Error; err != nil {
		return err
	}
	return nil
}

func FindHardDiskDrives() *[]models.MediaHardDrive {
	hdds := []models.MediaHardDrive{}
	db.Find(&hdds)
	return &hdds
}
