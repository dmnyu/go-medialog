package database

func FindResources() []Resource {
	resources := []Resource{}
	db.Find(&resources)
	return resources
}

func FindResource(id int) (Resource, error) {
	resource := Resource{}
	if err := db.Find(&resource, "id = ?", id).Error; err != nil {
		return resource, err
	}
	return resource, nil
}
