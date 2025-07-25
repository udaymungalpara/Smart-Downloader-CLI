package main

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/udaymungalpara/Smart-Downloader-CLI/internal/downloader"
)

func main() {
	source, _ := os.Open("./urls.txt")
	if source == nil {
		log.Fatal("No urls found in urls.txt")
	}
	var urls []string
	defer source.Close()
	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go downloader.Fetch(url, &wg)

	}
	wg.Wait()

}
