package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
    for _, url := range os.Args[1:] {
    	resp, err := http.Get(string(url))
    	if err != nil {
            fmt.Fprint(os.Stderr, "fetch:#{err}\n")
            os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch:reaading %s %v\n", url, err)
			os.Exit(1)
		}
	}
}