package index

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dmnyu/go-medialog/models"
	"github.com/dmnyu/go-medialog/shared"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
	"log"
	"strings"
)

const index = "media"

var (
	es      *elasticsearch7.Client
	indexes = []string{index}
	ctx     context.Context
)

func init() {
	es, _ = elasticsearch7.NewDefaultClient()
	ctx = context.Background()
}

func AddToIndex(entry models.MediaEntry) (string, error) {
	s, err := json.Marshal(entry)
	if err != nil {
		return "Could Not Marshal Entry", err
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

func DeleteFromIndex(docID string) error {
	resp, err := esapi.DeleteRequest{Index: index, DocumentID: docID}.Do(ctx, es.Transport)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("[INFO] %s", string(body))
	return nil
}

func FindDoc(docID string) (*models.MediaEntry, error) {

	resp, err := esapi.GetRequest{Index: index, DocumentID: docID}.Do(context.Background(), es.Transport)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	esHit := models.ESHit{}
	err = json.Unmarshal(body, &esHit)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] [INDEX] located document %s", docID)

	return &esHit.Source, nil

}

func SearchByAccessionID(accessionID int, pagination shared.Pagination) (*[]models.ESHit, error) {
	q := fmt.Sprintf(`{"query": {"match": {"accession_id": %d}}}`, accessionID)

	resp, err := esapi.SearchRequest{Index: indexes, Body: strings.NewReader(q)}.Do(context.Background(), es.Transport)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	esResponse := models.ESResponse{}
	err = json.Unmarshal(body, &esResponse)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] [INDEX] Located %d records for accession %d", len(esResponse.Hits.Hits), accessionID)

	return &esResponse.Hits.Hits, nil
}