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
	"strings"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

func getname(resp *http.Response, Url string) string {
	var filename string

	cd := resp.Header.Get("Content-Disposition")
	if cd != "" {
		rexp := regexp.MustCompile(`filename="(.+)"`)
		match := rexp.FindStringSubmatch(cd)
		if len(match) == 2 {
			return match[1]
		}
	}
	parsed, err := url.Parse(Url)
	if err != nil {
		return "download.tmp"
	}
	filename = path.Base(parsed.Path)
	if filename == "" || filename == "/" {
		return "download.tmp"

	}

	return filename

}

func retry(Url string, t time.Duration, attemp int) (*http.Response, error) {
	i := 0
	var err error
	for i < attemp {
		resp, err := http.Get(Url)
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}
		if resp != nil {
			resp.Body.Close()
			fmt.Printf("Retry %d for %s (Status: %d)\n", i+1, Url, resp.StatusCode)

		}

		time.Sleep(t)
		i++
	}
	return nil, fmt.Errorf("all retries failed %v", err)

}

func download(Url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := retry(Url, time.Second*1, 3)
	if resp == nil {
		fmt.Println("All retry are failed for:", err)
		return
	}
	defer resp.Body.Close()

	filename := getname(resp, Url)
	filename = strings.TrimSpace(filename)

	dst, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)

	_, err = io.Copy(io.MultiWriter(bar, dst), resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

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
