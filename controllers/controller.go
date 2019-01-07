package controllers

import (
	"github.com/gosexy/to"
	"reflect"
	"encoding/json"
	"fmt"
	"cherish-time-go/modules/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"log"
)

type Controller struct {
	beego.Controller
}

type RetData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *Controller) Prepare() {
}

//参数校验
func (c *Controller) Valid(inputData interface{}) {

	valid := validation.Validation{}
	b, err := valid.Valid(inputData)
	if err != nil {
		// handle error
	}
	if !b {
		// 处理抛出验证不通过
		for _, err := range valid.Errors {
			c.ThrowApi(-1, err.Key+" "+err.Message, "")
			log.Println(err.Key, err.Message)
		}
	}
}

func (c *Controller) ThrowApi(code int, msg string, data interface{}) {
	var retData RetData
	retData.Code = code
	retData.Msg = msg

	c.Data["json"] = retData
	c.ServeJSON()
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
