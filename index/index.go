package index

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dmnyu/go-medialog/database"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
	"log"
	"strings"
)

const index = "media"

var (
	es *elasticsearch7.Client
)

func init() {
	es, _ = elasticsearch7.NewDefaultClient()
}

func AddToIndex(entry database.MediaEntry) (string, error) {
	s, err := json.Marshal(entry)
	if err != nil {
		return "Could Not Marshall Entry", err
	}
	msg := bytes.NewReader(s)
	createRequest := esapi.IndexRequest{Index: index, Body: msg}
	resp, err := createRequest.Do(context.Background(), es.Transport)
	defer resp.Body.Close()
	if err != nil {
		return resp.String(), err
	}
	return resp.String(), nil
}

func SearchByAccessionID(accessionID int) ([]ESHits, error) {
	log.Println("[DEBUG] SEARCH BY ACCESSION CALLED")
	q := fmt.Sprintf(`{"query": {"match": {"accession_id": %d}}}`, accessionID)

	resp, err := esapi.SearchRequest{Index: []string{index}, Body: strings.NewReader(q)}.Do(context.Background(), es.Transport)
	if err != nil {
		return []ESHits{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println("[DEBUG]", string(body))
	if err != nil {
		return []ESHits{}, err
	}

	esResponse := ESResponse{}
	err = json.Unmarshal(body, &esResponse)
	if err != nil {
		return []ESHits{}, err
	}

	return esResponse.Hits.Hits, nil
}
