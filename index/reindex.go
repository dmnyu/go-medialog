package index

import (
	"github.com/dmnyu/go-medialog/database"
	"log"
)

func Reindex() error {
	//delete all from index
	/*
			q := `"{query": { "match_all": {}}}`
			resp, err := es.DeleteByQuery(indexes, strings.NewReader(q))
			if err != nil {
				panic(err)
			}


		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	*/

	//index each mediatype
	//optical disks
	log.Printf("[INFO] [INDEX] indexing optical discs")
	for _, disk := range *database.FindOpticaDiscs() {
		resp, err := AddToIndex(disk.GetMediaEntry())
		if err != nil {
			return err
		}
		log.Printf("[INFO] [INDEX] %s", resp)
	}

	//hard drives
	log.Printf("[INFO] [INDEX] indexing hard disk drives")
	for _, hdd := range *database.FindHardDiskDrives() {
		resp, err := AddToIndex(hdd.GetMediaEntry())
		if err != nil {
			return err
		}
		log.Printf("[INFO] [INDEX] %s", resp)
	}

	return nil
}
