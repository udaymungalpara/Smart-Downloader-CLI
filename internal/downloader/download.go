package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

func Download(Url string, wg *sync.WaitGroup) {

	defer wg.Done()
	var err error
	var resp *http.Response

	filename := Getname(Url) + ".part"
	final := strings.TrimSuffix(filename, ".part")
	if _, err := os.Stat(final); err == nil {
		fmt.Println("Already completed:", final)
		return
	}

	info, err := os.Stat(filename)
	if err == nil {
		offset := info.Size()
		err := Resume(Url, offset, filename)
		if err == nil {
			return
		}

	}
	resp, err = Retry(Url, time.Second*1, 3)

	if resp == nil {
		fmt.Println("All retry are failed for:", err)

		return
	}
	defer resp.Body.Close()
	dst, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)

	_, err = io.Copy(io.MultiWriter(bar, dst), resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	os.Rename(filename, strings.TrimSuffix(filename, ".part"))

}
