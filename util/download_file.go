package util

import (
	"io"
	"net/http"
	"os"
	"path"
)

// DownloadFile provides downloading of the bash script
// Its copy to the temp file
func DownloadFile(url string) (string, error) {
	client := &http.Client{}
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", nil
	}
	fileName := path.Base(r.URL.Path)
	out, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer out.Close()
	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
