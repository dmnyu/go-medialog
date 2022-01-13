package database

func FindResources() ([]Resource, error) {
	resources := []Resource{}
	if err := db.Find(&resources).Error; err != nil {
		return []Resource{}, err
	}
	return resources, nil
}

func FindResourcesByRepoID(id int) ([]Resource, error) {
	resources := []Resource{}
	err := db.Where("repository_id = ?", id).Find(&resources).Error
	return resources, err
}

func FindResource(id int) (Resource, error) {
	resource := Resource{}
	if err := db.Find(&resource, "id = ?", id).Error; err != nil {
		return resource, err
	}
	return resource, nil
}

func InsertResource(resource Resource) (int, error) {
	if err := db.Create(&resource).Error; err != nil {
		return 0, err
	}
	return int(resource.ID), nil
}

func GetNextMediaIDForResource(resourceID int) (int, error) {
	entry := MediaEntry{}
	if err := db.Order("media_id desc").Where("resource_id = ?", resourceID).First(&entry).Error; err != nil {
		return 0, err
	}

	return entry.MediaID + 1, nil
}
