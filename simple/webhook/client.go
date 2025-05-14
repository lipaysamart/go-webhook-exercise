package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var (
	httpClient = &http.Client{}
	webhook    = "http://localhost:8899/api/v1/wxwork/receive"
)

func main() {
	ctx := context.Background()

	payload := map[string]interface{}{
		"name": "hook",
		"data": map[string]interface{}{
			"text": "Hello, World!",
			"time": "2023-08-01 10:00:00",
		},
	}
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		log.Println("Encode Failed...", err)
		return
	}

	makeRequest(ctx, webhook, "POST", &buf)

}
func makeRequest(ctx context.Context, url string, method string, body io.Reader) error {
	httpReq, _ := http.NewRequestWithContext(ctx, method, url, body)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := httpClient.Do(httpReq)
	if httpResp.StatusCode != 200 || err != nil {
		log.Println("Send Request Failed...")
		return err
	}

	defer httpResp.Body.Close()
	return nil
}
