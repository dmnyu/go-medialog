package controllers

import "github.com/nyudlts/go-aspace"

var (
	client *aspace.ASClient
	err    error
)

func init() {
	client, err = aspace.NewClient("go-aspace.yml", "dev", 20)
	if err != nil {
		panic(err)
	}
}

func FindAspaceResource(repositoryID int, resourceID int) (aspace.Resource, error) {
	resource, err := client.GetResource(repositoryID, resourceID)
	if err != nil {
		return resource, err
	}
	return resource, nil
}

func FindAspaceAccession(repositoryID int, accessionID int) (aspace.Accession, error) {
	accession, err := client.GetAccession(repositoryID, accessionID)
	if err != nil {
		return accession, err
	}
	return accession, nil
}
