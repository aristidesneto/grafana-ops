package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type Client struct {
	baseURL string
	token   string
	logger  *logrus.Logger
}

func NewClient(url, token string, log *logrus.Logger) *Client {
	if url == "" {
		log.Error("Grafana Url not found, please provide")
		os.Exit(1)
	}

	if token == "" {
		log.Error("Grafana Token not found, please provide")
		os.Exit(1)
	}

	return &Client{
		baseURL: url,
		token:   token,
		logger:  log,
	}
}

func (c *Client) doRequest(method, path string, body ...[]byte) ([]byte, error) {
	url := c.baseURL + path
	var reqBody io.Reader
	if len(body) > 0 {
		reqBody = bytes.NewReader(body[0])
	}

	c.logger.Debugf("%s %s", method, url)
	if reqBody != nil {
		c.logger.Debugf("request body size %d", len(body[0]))
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error processing request: %w", err)
	}
	defer resp.Body.Close()

	c.logger.Debugf("response status %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error in api response (status %d): %s", resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
}

func (c *Client) Get(path string) ([]byte, error) {
	return c.doRequest("GET", path)
}

func (c *Client) Post(path string, body []byte) ([]byte, error) {
	return c.doRequest("POST", path, body)
}

func (c *Client) Put(path string, body []byte) ([]byte, error) {
	return c.doRequest("PUT", path, body)
}

func (c *Client) Delete(path string) ([]byte, error) {
	return c.doRequest("DELETE", path)
}
