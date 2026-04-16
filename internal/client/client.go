package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const headerAPIKey = "X-API-Key"

type Client struct {
	baseURL    *url.URL
	apiKey     string
	httpClient *http.Client
}

type APIStatusError struct {
	Status int
	Body   []byte
}

func New(rawBaseURL, apiKey string, timeout time.Duration) (*Client, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, fmt.Errorf("api key is empty")
	}
	raw := strings.TrimSpace(rawBaseURL)
	if raw == "" {
		raw = "https://api.wallbit.io"
	}
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, fmt.Errorf("invalid base URL %q", rawBaseURL)
	}
	u.RawQuery = ""
	u.Fragment = ""
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	return &Client{
		baseURL: u,
		apiKey:  strings.TrimSpace(apiKey),
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c *Client) Get(ctx context.Context, path string) ([]byte, int, error) {
	if !strings.HasPrefix(path, "/") {
		return nil, 0, fmt.Errorf("path must start with /, got %q", path)
	}
	rel, err := url.Parse(path)
	if err != nil {
		return nil, 0, fmt.Errorf("parse path: %w", err)
	}
	full := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, full.String(), nil)
	if err != nil {
		return nil, 0, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set(headerAPIKey, c.apiKey)
	req.Header.Set("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("http get: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(io.LimitReader(res.Body, 1<<20))
	if err != nil {
		return nil, res.StatusCode, fmt.Errorf("read body: %w", err)
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return body, res.StatusCode, APIStatusError{Status: res.StatusCode, Body: body}
	}

	return body, res.StatusCode, nil
}

func (e APIStatusError) Error() string {
	msg, _ := parseErrorMessage(e.Body)
	if msg != "" {
		return fmt.Sprintf("wallbit API %d: %s", e.Status, msg)
	}
	return fmt.Sprintf("wallbit API %d", e.Status)
}

func parseErrorMessage(body []byte) (string, error) {
	var wrap struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &wrap); err != nil || wrap.Message == "" {
		return "", err
	}
	return wrap.Message, nil
}
