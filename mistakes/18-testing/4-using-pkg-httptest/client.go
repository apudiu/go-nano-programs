package tst

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type DurationClient struct {
	client *http.Client
}

func (c DurationClient) GetDuration(url string, lat1, lng1, lat2, lng2 float64) (
	time.Duration, error,
) {
	resp, err := c.client.Post(
		url, "application/json",
		buildRequestBody(lat1, lng1, lat2, lng2),
	)
	if err != nil {
		return 0, err
	}

	return parseResponseBody(resp.Body)
}

func NewDurationClient() DurationClient {
	return DurationClient{
		client: http.DefaultClient,
	}
}

func buildRequestBody(lat1, lng1, lat2, lng2 float64) io.Reader {
	return strings.NewReader("")
}

type request struct {
	Duration int
}

func parseResponseBody(r io.ReadCloser) (time.Duration, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = r.Close()
	}()

	var req request
	err = json.Unmarshal(b, &req)
	if err != nil {
		return 0, err
	}
	return time.Duration(req.Duration) * time.Second, nil
}
