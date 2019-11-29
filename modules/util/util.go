package util

import (
	uuid "github.com/satori/go.uuid"
	"encoding/json"
	"github.com/astaxie/beego/context"
	"strings"
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
	uuidFunc := uuid.NewV4()
	uuidStr := uuidFunc.String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	uuidByt := []rune(uuidStr)
	return string(uuidByt[8:24])
}

func DaysDiff(startDate, endDate int64) int64 {
	if startDate > endDate {
		return -1
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
