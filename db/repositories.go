package db

func FindRepository(id int) (Repository, error) {
	repository := Repository{}
	if err := DB.Find(&repository, "id = ?", id).Error; err != nil {
		return repository, err
	}

	return repository, nil
}
