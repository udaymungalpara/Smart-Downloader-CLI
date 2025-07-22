package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"sync"
)

func download(Url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(Url)
	if err != nil {
		fmt.Println(err)
	}
	var filename string
	cd := resp.Header.Get("Content-Disposition")
	if cd != "" {
		rexp := regexp.MustCompile(`filename="(.+)"`)
		match := rexp.FindStringSubmatch(cd)
		if len(match) == 2 {
			filename = match[1]
		}

	} else {
		parsed, err := url.Parse(Url)
		if err != nil {
			filename = "download.tmp"
		} else {
			filename = path.Base(parsed.Path)
			if filename == "" || filename == "/" {
				filename = "download.tmp"
			}
		}
	}

	dst, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	_, err = io.Copy(dst, resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

}

func main() {
	source, _ := os.Open("urls.txt")
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

		go download(url, &wg)

	}
	wg.Wait()

}
