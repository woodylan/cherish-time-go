package util

import (
	"encoding/json"
	"strings"
	"github.com/astaxie/beego/context"
	"github.com/satori/go.uuid"
	"strconv"
	"fmt"
)

type RetData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}



func JsonDecode(jsonStr string, structModel interface{}) error {
	decode := json.NewDecoder(strings.NewReader(jsonStr))
	err := decode.Decode(structModel)
	return err
}

func JsonEncode(structModel interface{}) (string, error) {
	jsonStr, err := json.Marshal(structModel)
	return string(jsonStr), err
}

func GenShortUuid() string {
	dict := map[int]byte{
		0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9', 10: 'a', 11: 'b', 12: 'c',
		13: 'd', 14: 'e', 15: 'f', 16: 'g', 17: 'h', 18: 'i', 19: 'j', 20: 'k', 21: 'l', 22: 'm', 23: 'n', 24: 'o',
		25: 'p', 26: 'q', 27: 'r', 28: 's', 29: 't', 30: 'u', 31: 'v', 32: 'w', 33: 'x', 34: 'y', 35: 'z', 36: 'A',
		37: 'B', 38: 'C', 39: 'D', 40: 'E', 41: 'F', 42: 'G', 43: 'H', 44: 'I', 45: 'J', 46: 'K', 47: 'L', 48: 'M',
		49: 'N', 50: 'O', 51: 'P', 52: 'Q', 53: 'R', 54: 'S', 55: 'T', 56: 'U', 57: 'V', 58: 'W', 59: 'X', 60: 'Y',
		61: 'Z', 62: '0', 63: '1'}

	uuidSrcStr := uuid.NewV4().String()
	uuidStr := strings.Replace(uuidSrcStr+"0000", "-", "", -1)

	var shortUuid string

	for index := 0; index < 6; index++ {
		word, _ := strconv.ParseInt(string(uuidStr[index*6:index*6+6]), 16, 0)
		binNum := fmt.Sprintf("%024b", word)
		for bit := 0; bit < 4; bit++ {
			stepBit := binNum[bit*6 : bit*6+6]
			dexNum, _ := strconv.ParseInt(stepBit, 2, 0)
			if index == 5 && bit >= 2 {
				continue
			}
			shortUuid += string(dict[int(dexNum)])
		}
	}
	return shortUuid
}

func DaysDiff(startDate, endDate int64) int64 {
	if startDate > endDate {
		return 0
	}

	betweenTime := endDate - startDate

	return betweenTime / 86400
}

func ThrowApi(ctx *context.Context, code int, msg string) (*context.Context) {
	var retData RetData
	retData.Code = code
	retData.Msg = msg

	ctx.Output.JSON(retData, true, false)
	return ctx
}
