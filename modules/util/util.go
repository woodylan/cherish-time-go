package util

import (
	"encoding/json"
	"strings"
	)

func JsonDecode(jsonStr string, structModel interface{}) error {
	decode := json.NewDecoder(strings.NewReader(jsonStr))
	err := decode.Decode(structModel)
	return err
}

func DaysDiff(startDate, endDate int64) int64 {
	if startDate > endDate {
		return 0
	}

	betweenTime := endDate - startDate

	return betweenTime / 86400
}
