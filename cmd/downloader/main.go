package main

import (
	"bufio"
	"log"
	"os"
	"sync"
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

		go downloader.fetch(url)

	}
	wg.Wait()

}
