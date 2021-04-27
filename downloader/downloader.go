package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

//DownloadFile Downloads file
func DownloadFile(url string) (bool, error) {
	res, err := http.Head(url)

	if err != nil {
		return false, err
	}

	header := res.Header
	contentLength, _ := strconv.Atoi(header["Content-Length"][0])
	fmt.Println(contentLength)
	threads := 10
	fileLengths := contentLength / threads
	diff := contentLength % threads

	for i := 0; i < threads; i++ {
		wg.Add(1)

		min := fileLengths * i
		max := fileLengths * (i + 1)

		if i == threads-1 {
			max += diff
		}

		go func(min int, max int, i int) {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", url, nil)
			rangeHeader := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1)
			req.Header.Add("range", rangeHeader)
			resp, _ := client.Do(req)
			defer resp.Body.Close()
			reader, _ := ioutil.ReadAll(resp.Body)
			ioutil.WriteFile(strconv.Itoa(i), reader, 0x777)
			wg.Done()
		}(min, max, i)
	}
	wg.Wait()
	return true, nil
}
