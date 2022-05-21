package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type result struct {
	Args string `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin string `json:"origin"`
	Url string `json:"url"`
}

func main() {
	resp, err := http.Get("https://api.64clouds.com/v1/getLiveServiceInfo?veid=835634&api_key=API_KEY")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	var res result
	_ = json.Unmarshal(body,&res)
	fmt.Printf("%#v", res)
}

//$request = "https://api.64clouds.com/v1/getServiceInfo?veid=835634&api_key=YOUR_API_KEY_HERE";
//$serviceInfo = json_decode (file_get_contents ($request));
//print_r ($serviceInfo);
