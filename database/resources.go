package database

func FindResources() []Resource {
	resources := []Resource{}
	db.Find(&resources)
	return resources
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
