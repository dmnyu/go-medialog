package database

func FindResources() ([]Resource, error) {
	resources := []Resource{}
	if err := db.Find(&resources).Error; err != nil {
		return []Resource{}, err
	}
	return resources, nil
}

func FindResourcesByRepoID(id int, pagination Pagination) ([]Resource, error) {
	resources := []Resource{}
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	err := queryBuider.Where("repository_id = ?", id).Find(&resources).Error
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
	entry := []MediaEntry{}
	if err := db.Order("media_id desc").Select("media_id").Where(&entry, "resource_id = ?", resourceID).Error; err != nil {
		return 0, err
	}

	if len(entry) == 0 {
		return 1, nil
	}

	return entry[0].MediaID + 1, nil
}
