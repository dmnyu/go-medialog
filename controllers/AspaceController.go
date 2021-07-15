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

func getAspaceAccession(repositoryID int, accessionID int) (aspace.Accession, error) {
	accession, err := client.GetAccession(repositoryID, accessionID)
	if err != nil {
		return accession, err
	}
	return accession, nil
}

func getAccessionIndentifierString(a aspace.Accession) string {
	s := a.ID0;
	for _, id := range[]string{a.ID1, a.ID2, a.ID3} {
		if id != "" {
			s = s + "." + id
		}
	}
	return s
}
