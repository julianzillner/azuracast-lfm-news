package download

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func Download(url string) error {
	out, err := os.Create("news.mp3")
	if err != nil {
		return err
	}
	defer out.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if strings.Contains(url, "@") {
		parts := strings.Split(url, "@")
		authPart := strings.Split(strings.TrimPrefix(parts[0], "https://"), ":")
		if len(authPart) == 2 {
			req.SetBasicAuth(authPart[0], authPart[1])
			req.URL.Scheme = "https"
			req.URL.Host = strings.Split(parts[1], "/")[0]
			req.URL.Path = "/" + strings.Join(strings.Split(parts[1], "/")[1:], "/")
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	return err
}
