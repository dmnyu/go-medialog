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
