package downloader

import (
	"fmt"
	"net/http"
	"time"
)

func Retry(Url string, t time.Duration, attemp int) (*http.Response, error) {
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
