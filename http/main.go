package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWwriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWwriter{}

	io.Copy(lw, resp.Body)
}

func (logWwriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
