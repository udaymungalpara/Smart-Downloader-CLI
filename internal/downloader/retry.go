package downloader

import (
	"fmt"
	"net/http"
	"time"
)

func Retry(url string, delay time.Duration, attempts int) (*http.Response, error) {
	var lasterr error

	for i := 0; i < attempts; i++ {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}
		if resp != nil {
			resp.Body.Close()
			lasterr = fmt.Errorf("retry %d for %s (Status: %d)", i+1, url, resp.StatusCode)

		} else {
			lasterr = err
		}

		time.Sleep(delay)
	}
	return nil, fmt.Errorf("all retries failed %v", lasterr)

}
