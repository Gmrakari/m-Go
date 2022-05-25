package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
)

func main() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["meg"] = "success"
	res["data"] = map[string]interface{} {
		"username" : "Tom",
		"age"      : "30",
		"hobby"    : []string{"reading", "playing"},
	}
	fmt.Println("map data:", res)

	//json 序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("----map to json ---")
	fmt.Println("json data :", string(jsons))

    //反序列化
    res2 := make(map[string]interface{})
    errs = json.Unmarshal([]byte(jsons), &res2)
    if errs != nil {
    	fmt.Println("Json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("----map to json ---")
    fmt.Println("map data:", res2)

}
