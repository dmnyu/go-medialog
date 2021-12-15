package database

func FindAccessions() []Accession {
	accessions := []Accession{}
	db.Find(&accessions)
	return accessions
}

func FindAccession(id int) (Accession, error) {
	accession := Accession{}
	if err := db.Where("id = ?", id).First(&accession).Error; err != nil {
		return accession, err
	}

	return accession, nil
}

func InsertAccession(accession Accession) (int, error) {
	if err := db.Create(&accession).Error; err != nil {
		return 0, err
	}

	return int(accession.ID), nil
}
