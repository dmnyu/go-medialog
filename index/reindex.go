package index

import (
	"github.com/dmnyu/go-medialog/database"
	"log"
)

func Reindex() error {

	//index each mediatype
	//optical disks
	log.Printf("[INFO] [INDEX] indexing optical discs")
	for _, disk := range *database.FindOpticaDiscs() {
		resp, err := AddToIndex(disk.GetMediaEntry(), nil)
		if err != nil {
			return err
		}
		log.Printf("[INFO] [INDEX] %s", resp)
	}

	//hard drives
	log.Printf("[INFO] [INDEX] indexing hard disk drives")
	for _, hdd := range *database.FindHardDiskDrives() {
		resp, err := AddToIndex(hdd.GetMediaEntry(), nil)
		if err != nil {
			return err
		}
		log.Printf("[INFO] [INDEX] %s", resp)
	}

	return nil
}
