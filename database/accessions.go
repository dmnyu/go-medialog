package database

import "github.com/dmnyu/go-medialog/models"

func FindAccessions() []models.Accession {
	accessions := []models.Accession{}
	db.Find(&accessions)
	return accessions
}

func FindAccession(id int) (models.Accession, error) {
	accession := models.Accession{}
	if err := db.Where("id = ?", id).First(&accession).Error; err != nil {
		return accession, err
	}

	return accession, nil
}

func InsertAccession(accession models.Accession) (int, error) {
	if err := db.Create(&accession).Error; err != nil {
		return 0, err
	}

	return int(accession.ID), nil
}
