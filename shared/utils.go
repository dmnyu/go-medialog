package shared

import "os"
import "errors"

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func FileExists(location string) error {
	if _, err := os.Stat(location); err == nil {
		return nil
	} else if errors.Is(err, os.ErrNotExist) {
		return err
	} else {
		return err
	}
}
