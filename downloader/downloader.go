package downloader

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func GetTextFromUrl(url string, c chan string) error {
	resp, err := http.Get(url)

	fmt.Println("Downloading from: ", url)

	if err != nil {
		fmt.Println(resp, err)
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return err
	}

	c <- string(body)
	return nil
}

func GetTextParallel(urls []string) []string {
	results := make([]string, len(urls))

	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(urls))

	for i := 0; i < len(urls); i++ {
		defer waitGroup.Done()
		c := make(chan string)
		go GetTextFromUrl(urls[i], c)
		res := <-c
		results[i] = res
	}

	return results
}
