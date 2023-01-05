package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Fetchall fetches URLs in parallel and reports their times and sizes.

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {  // 启动一个goroutine
		go fetch_2(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)              // 从通道ch接收
	}

	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch_2(url string, ch chan<- string) {
	i, j := strings.Index(url, "."), strings.LastIndex(url, ".")

	// in case of only one "." in url

	if i == j {
		i = 0
	}
	output, err := os.Create(url[i+1:j] + "-dump.html")

	start := time.Now()
	if err != nil {
		ch <- fmt.Sprintf("while create output file:%v", err)
		return
	}

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(output, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

