package util

import (
	"encoding/json"
	"strings"
	"github.com/astaxie/beego/context"
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
