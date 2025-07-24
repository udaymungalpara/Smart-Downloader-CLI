package downloader

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
)

func Getname(Url string) string {

	req, err := http.NewRequest("HEAD", Url, nil)
	if err != nil {
		fmt.Println("Initial GET failed for filename extraction:", err)
		return "download.tmp.part"
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

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
	filename := path.Base(parsed.Path)
	if filename == "" || filename == "/" {
		return "download.tmp"

	}

	return filename

}
