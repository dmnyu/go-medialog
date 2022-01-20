package database

import "github.com/dmnyu/go-medialog/models"

func InsertOpticalDisc(disc *models.MediaOpticalDisc) error {
	if err := db.Create(disc).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOpticalDisc(discID int) error {
	if err := db.Delete(&models.MediaOpticalDisc{}, discID).Error; err != nil {
		return err
	}
	return nil
}
