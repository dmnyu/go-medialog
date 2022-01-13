package database

func InsertEntry(entry *MediaEntry) error {
	if err := db.Create(&entry).Error; err != nil {
		return err
	}
	return nil
}
