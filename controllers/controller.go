package controllers

import (
	"encoding/json"
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

type Page struct {
	Count       int         `json:"count"`
	PerPage     int         `json:"perPage"`
	CurrentPage int         `json:"currentPage"`
	LastPage    int         `json:"lastPage"`
	List        interface{} `json:"list"`
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
			util.ThrowApi(c.Ctx, -1, err.Key+" "+err.Message)
			log.Println(err.Key, err.Message)
		}
	}
}

func (c *Controller) GetData(RequestData interface{}) {
	data := c.GetString("data")

	//string 转 struct
	if err := json.Unmarshal([]byte(data), &RequestData); err == nil {
	}
}

func (page *Page) RendPage(count, perPage, currentPage int) (*Page) {
	lastPage := count / perPage
	if count%perPage > 0 {
		lastPage += 1
	}

	page.Count = count
	page.PerPage = perPage
	page.CurrentPage = currentPage
	page.LastPage = lastPage

	return page
}
