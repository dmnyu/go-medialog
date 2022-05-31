package controllers

import "github.com/nyudlts/go-aspace"

var (
	client       *aspace.ASClient
	err          error
	AspaceEnv    string
	AspaceConfig string
)

func SetEnvironment(env string) {
	AspaceEnv = env
}

func GetClient() error {
	client, err = aspace.NewClient("config/go-aspace.yml", AspaceEnv, 20)
	if err != nil {
		return err
	}
	return nil
}

func FindAspaceRepository(repositoryID int) (aspace.Repository, error) {
	GetClient()
	repository, err := client.GetRepository(repositoryID)
	return repository, err
}

func FindAspaceResource(repositoryID int, resourceID int) (aspace.Resource, error) {
	GetClient()
	resource, err := client.GetResource(repositoryID, resourceID)
	if err != nil {
		return resource, err
	}
	return resource, nil
}

func FindAspaceAccession(repositoryID int, accessionID int) (aspace.Accession, error) {
	GetClient()
	accession, err := client.GetAccession(repositoryID, accessionID)
	if err != nil {
		return accession, err
	}
	return accession, nil
}

func GetResourceList(repositoryID int) ([]aspace.ResourceListEntry, error) {
	GetClient()
	return client.GetResourceList(repositoryID)
}

func GetAccessionListForResource(repositoryID int, resourceID int) ([]aspace.AccessionEntry, error) {
	if client == nil {
		GetClient()
	}
	return client.GetAccessionList(repositoryID, resourceID)
}

func GetASpaceRepositories() (*map[int]string, error) {
	GetClient()
	repositoryIDs, err := client.GetRepositories()
	if err != nil {
		return nil, err
	}

	repositories := map[int]string{}
	for _, repoID := range repositoryIDs {
		repository, err := client.GetRepository(repoID)
		if err != nil {
			return nil, err
		}
		repositories[repoID] = repository.Name
	}

	return &repositories, nil
}
