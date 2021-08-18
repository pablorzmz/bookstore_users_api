package date_utils

import "time"

const apiDateLayaut = "2006-01-02T15:04:05Z"

func GetNowString() string {
	return time.Now().UTC().Format(apiDateLayaut)
}
