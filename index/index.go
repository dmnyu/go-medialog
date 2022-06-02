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
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"io/ioutil"
	"log"
	"strings"
)

const index = "media"

var (
	es        *elasticsearch7.Client
	indexes   = []string{index}
	ctx       context.Context
	transport estransport.Interface
)

func init() {
	es, _ = elasticsearch7.NewDefaultClient()
	ctx = context.Background()
	transport = es.Transport
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

func UpdateDocument(entry models.MediaEntry, docID string) (*models.ESCreateResponse, error) {

	if err := DeleteFromIndex(docID); err != nil {
		return nil, err
	}

	msg, err := AddToIndex(entry)
	if err != nil {
		return nil, err
	}

	response := models.ESCreateResponse{}
	json.Unmarshal([]byte(msg), &response)

	return &response, nil
}

func DeleteFromIndex(docID string) error {
	_, err := esapi.DeleteRequest{Index: index, DocumentID: docID}.Do(ctx, es.Transport)
	if err != nil {
		return err
	}
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

	return &esResponse.Hits.Hits, nil
}

func FindNextMediaIDInResource(resourceID int) (*int, error) {
	//construct query
	q := fmt.Sprintf(`{"query": {"match": {"resource_id": %d}}}`, resourceID)

	//make request
	resp, err := esapi.SearchRequest{Index: indexes, Body: strings.NewReader(q)}.Do(context.Background(), es.Transport)
	if err != nil {
		return nil, err
	}

	//get response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//Unmarshal response
	esResponse := models.ESResponse{}
	if err := json.Unmarshal(body, &esResponse); err != nil {
		return nil, err
	}

	//get the next mediaID for the resource
	nextMediaId := 0
	for _, hit := range esResponse.Hits.Hits {
		if hit.Source.MediaID > nextMediaId {
			nextMediaId = hit.Source.MediaID
		}
	}
	nextMediaId++
	return &nextMediaId, nil
}

func KeywordSearch(query string) (*[]models.ESHit, error) {
	//format the query
	q := fmt.Sprintf(`{"query":{"match":{"json":{"query":"%s"}}}}`, query)

	//make request
	resp, err := esapi.SearchRequest{Index: indexes, Body: strings.NewReader(q)}.Do(context.Background(), es.Transport)
	if err != nil {
		return nil, err
	}

	//get response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//Unmarshal response
	esResponse := models.ESResponse{}
	err = json.Unmarshal(body, &esResponse)
	if err != nil {
		return nil, err
	}

	return &esResponse.Hits.Hits, nil
}

func DeleteAll() error {

	var buf = bytes.Buffer{}
	// Elastic query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return err
	}

	body := strings.NewReader(buf.String())

	resp, err := es.DeleteByQuery(indexes, body)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("[INFO] [INDEX] %s", string(respBody))
	return nil
}
