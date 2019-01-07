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
