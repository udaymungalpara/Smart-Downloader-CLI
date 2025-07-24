package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func Resume(Url string, offset int64, filename string) (err error) {
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-", offset))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"resuming",
	)

	write, errr := io.Copy(io.MultiWriter(bar, file), resp.Body)
	defer func() {
		file.Close() 
		if errr == nil && write == resp.ContentLength {
			final := strings.TrimSuffix(filename, ".part")
			if err := os.Rename(filename, final); err != nil {
				fmt.Println("Rename fail:", err)
			} else {
				fmt.Println("Resumed and saved:", final)
			}
		}
	}()

	return errr

}
