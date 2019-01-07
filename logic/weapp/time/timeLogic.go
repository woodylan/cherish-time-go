package timeLogic

import (
	"cherish-time-go/models/Time"
	"fmt"
)

func GetDetail(id string) {
	logic := TimeModel.GetById("10ab11e9840d9b58")

	fmt.Println(logic.Color)
}
