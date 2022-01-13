package database

func InsertOpticalDisc(disc *MediaOpticalDisc) error {
	if err := db.Create(disc).Error; err != nil {
		return err
	}
	return nil
}
