package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	namefile := os.Args[1]
	file, err := os.Open(namefile)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
