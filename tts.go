package apiai

import (
	"hash/fnv"
	"io"
	"net/http"
	"os"
	"strconv"
)

func (c *ApiClient) Tts(text string) (string, error) {

	req, err := http.NewRequest("GET", c.buildUrl("tts", map[string]string{
		"text": text,
	}), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.config.Token)
	req.Header.Set("Accept-Language", c.config.SpeechLang)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	filePath := "/tmp/" + hash(text)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	io.Copy(file, resp.Body)

	return filePath, nil
}

func hash(s string) string {
	h := fnv.New64a()
	h.Write([]byte(s))
	hValue := strconv.FormatUint(h.Sum64(), 10)
	return hValue
}
