package db

func FindResource(resourceID int) (Resource, error) { // Get model if exist
	var resource Resource

	if err := DB.Where("id = ?", resourceID).First(&resource).Error; err != nil {
		return resource, err
	}
	return resource, nil
}
