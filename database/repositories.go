package database

func FindRepositories() []Repository {
	repositories := []Repository{}
	db.Find(&repositories)
	return repositories
}

func FindRepository(id int) (Repository, error) {
	repository := Repository{}
	if err := db.Find(&repository, "id = ?", id).Error; err != nil {
		return repository, err
	}
	return repository, nil
}

func InsertRepository(repo Repository) error {
	if err := db.Create(&repo).Error; err != nil {
		return err
	}
	return nil
}

/*
func DeleteRepository(id int) (int, error) {
	// Get model if exist
	var repository Repository
	if err := db.Where("id = ?",id).First(&repository).Error; err != nil {
		return 0, err
	}

	if err := db.Delete(&repository).Error; err != nil {
		return 0, err
	}

	return id, nil
}
*/
