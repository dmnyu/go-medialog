package database

import (
	"github.com/dmnyu/go-medialog/models"
	"log"
)

func InsertOpticalDisc(disc *models.MediaOpticalDisc) error {
	if err := db.Create(disc).Error; err != nil {
		return err
	}
	log.Printf("[INFO] created optical disc with id %d", disc.ID)
	return nil
}

func DeleteOpticalDisc(discID uint) error {
	if err := db.Delete(&models.MediaOpticalDisc{}, discID).Error; err != nil {
		return err
	}
	log.Printf("[INFO] deleted optical disc with id %d", discID)
	return nil
}

func FindOpticalDisc(discID int) (models.MediaOpticalDisc, error) {
	o := models.MediaOpticalDisc{}
	if err := db.Find(&o, discID).Error; err != nil {
		return o, err
	}
	return o, nil
}
