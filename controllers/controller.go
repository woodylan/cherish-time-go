package controllers

import (
	"github.com/gosexy/to"
	"reflect"
	"encoding/json"
	"fmt"
	"cherish-time-go/modules/util"
	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) GetData(RequestData interface{}) {
	s := reflect.ValueOf(RequestData).Elem()
	typeOf := s.Type()

	for i := 0; i < s.NumField(); i++ {

		data := c.GetString("data")

		//string 转 map
		var dataMap map[string]interface{}
		if err := json.Unmarshal([]byte(data), &dataMap); err == nil {
		}

		field := typeOf.Field(i)
		key := field.Tag.Get("data") //取Tag里的字段

		//todo 出错判断
		dataValue := dataMap[key]

		switch fmt.Sprintf("%s", field.Type) {
		case "string":
			s.Field(i).SetString(to.String(dataValue))
		case "int":
			fallthrough
		case "int8":
			fallthrough
		case "int16":
			fallthrough
		case "int32":
			fallthrough
		case "int64":
			s.Field(i).SetInt(to.Int64(dataValue))
		case "uint":
			fallthrough
		case "uint8":
			fallthrough
		case "uint16":
			fallthrough
		case "uint32":
			fallthrough
		case "uint64":
			s.Field(i).SetUint(to.Uint64(dataValue))
		case "[]string":
			tmpSlice := make([]string, 0)
			util.JsonDecode(data, &tmpSlice)
			s.Field(i).Set(reflect.ValueOf(tmpSlice))
		case "[]uint8":
			tmpSlice := make([]uint8, 0)
			util.JsonDecode(data, &tmpSlice)
			s.Field(i).Set(reflect.ValueOf(tmpSlice))
		case "[]int":
			tmpSlice := make([]int, 0)
			util.JsonDecode(data, &tmpSlice)
			s.Field(i).Set(reflect.ValueOf(tmpSlice))
		case "[]int32":
			tmpSlice := make([]int32, 0)
			util.JsonDecode(data, &tmpSlice)
			s.Field(i).Set(reflect.ValueOf(tmpSlice))
		case "[]int64":
			tmpSlice := make([]int64, 0)
			util.JsonDecode(data, &tmpSlice)
			s.Field(i).Set(reflect.ValueOf(tmpSlice))
		}
	}
}