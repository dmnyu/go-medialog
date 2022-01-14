package database

import "log"

func InsertEntry(entry *MediaEntry) error {
	if err := db.Create(&entry).Error; err != nil {
		return err
	}
	return nil
}

func DeleteEntry(entryID int) error {
	if err := db.Delete(&MediaEntry{}, entryID).Error; err != nil {
		return err
	}
	return nil
}

func FindEntry(entryID int) (MediaEntry, error) {
	log.Println(entryID)
	entry := MediaEntry{}
	if err := db.Where("id = ?", entryID).Find(&entry).Error; err != nil {
		return MediaEntry{}, err
	}

	log.Println(entry.ID)

	return entry, nil
}

func FindByAccessionID(accessionID int) ([]MediaEntry, error) {
	entries := []MediaEntry{}
	err := db.Where("accession_id = ?", accessionID).Find(&entries).Error
	if err != nil {
		return entries, err
	}
	return entries, nil
}
