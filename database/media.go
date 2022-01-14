package database

func InsertOpticalDisc(disc *MediaOpticalDisc) error {
	if err := db.Create(disc).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOpticalDisc(discID int) error {
	if err := db.Delete(&MediaOpticalDisc{}, discID).Error; err != nil {
		return err
	}
	return nil
}
