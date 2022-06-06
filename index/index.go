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

type ObjectID int

const (
	Repository ObjectID = iota
	Resource
	Accession
)

var ObjectIds = map[ObjectID]string{
	Repository: "repository_id",
	Resource:   "resource_id",
	Accession:  "accession_id",
}

func init() {
	es, _ = elasticsearch7.NewDefaultClient()
	ctx = context.Background()
	transport = es.Transport
}

func AddToIndex(entry models.MediaEntry, docID *string) (*models.ESTXResponse, error) {
	createEntry, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}

	createRequest := esapi.IndexRequest{Index: index, Body: bytes.NewReader(createEntry)}
	if docID != nil {
		createRequest.DocumentID = *docID
	}

	resp, err := createRequest.Do(ctx, transport)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	createResponse := models.ESTXResponse{}
	if err := json.Unmarshal(body, &createResponse); err != nil {
		return nil, err
	}
	createResponse.Json = string(body)

	return &createResponse, nil
}

func DeleteFromIndex(docID string) (*models.ESTXResponse, error) {
	deleteRequest, err := esapi.DeleteRequest{Index: index, DocumentID: docID}.Do(ctx, es.Transport)
	if err != nil {
		return nil, err
	}
	defer deleteRequest.Body.Close()
	body, err := ioutil.ReadAll(deleteRequest.Body)
	if err != nil {
		return nil, err
	}

	deleteResponse := models.ESTXResponse{}
	if err := json.Unmarshal(body, &deleteResponse); err != nil {
		return nil, err
	}
	deleteResponse.Json = string(body)
	return &deleteResponse, nil
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

func FindByType(id int, objectTypeID ObjectID, pagination shared.Pagination) (*[]models.ESHit, error) {

	q := fmt.Sprintf(`{"query": {"match": {"%s": %d}}}`, ObjectIds[objectTypeID], id)
	log.Printf("%v", q)
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
