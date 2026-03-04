package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// generic HTTP GET wrapper.
func FetchandDecode[T any](ctx context.Context, url string, headers map[string]string) (*T, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	// attach provided headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// timeout check for the get request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected statud code: %d", resp.StatusCode)
	}

	// decode json response into a generic type T
	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
