package database

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/dmnyu/go-medialog/shared"
)

func FindResources() ([]models.Resource, error) {
	resources := []models.Resource{}
	if err := db.Find(&resources).Error; err != nil {
		return []models.Resource{}, err
	}
	return resources, nil
}

func FindResourcesByRepoID(id int, pagination shared.Pagination) ([]models.Resource, error) {
	resources := []models.Resource{}
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	err := queryBuider.Where("repository_id = ?", id).Find(&resources).Error
	return resources, err
}

func FindResource(id int) (models.Resource, error) {
	resource := models.Resource{}
	if err := db.Find(&resource, "id = ?", id).Error; err != nil {
		return resource, err
	}
	return resource, nil
}

func InsertResource(resource models.Resource) (int, error) {
	if err := db.Create(&resource).Error; err != nil {
		return 0, err
	}
	return int(resource.ID), nil
}

/* this is broken, get the next media id from the index */
func GetNextMediaIDForResource(resourceID int) (int, error) {
	entry := []models.MediaEntry{}
	if err := db.Order("media_id desc").Select("media_id").Where(&entry, "resource_id = ?", resourceID).Error; err != nil {
		return 0, err
	}

	if len(entry) == 0 {
		return 1, nil
	}

	return entry[0].MediaID + 1, nil
}

func DeleteResource(resource *models.Resource) error {
	if err := db.Delete(resource).Error; err != nil {
		return err
	}
	return nil
}
