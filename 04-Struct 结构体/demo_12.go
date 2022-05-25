package main

// 改变数据
import (
	"encoding/json"
	"fmt"
)

type M_Result struct {
	Code int `json:"code"`
	Message string `json:"msg"`
}

func main() {
	var res M_Result
	res.Code = 200
	res.Message = "success"
    toJson(&res)

	setData(&res)
	toJson(&res)
}

func setData(res *M_Result) {
	res.Code = 500
	res.Message = "Server Error"
}

func toJson(res *M_Result) {
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))
}
