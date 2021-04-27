package main

import (
	"fmt"
	"os"

	"example.com/downloader"
)

func main() {
	url := os.Args[1]
	x, err := downloader.DownloadFile(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(x)
}
