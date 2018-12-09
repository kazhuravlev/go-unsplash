package unsplash

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	log        *logrus.Logger
}

type Option func(*Client) error

func WithHttpClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}

func WithLogrus(logger *logrus.Logger) Option {
	return func(c *Client) error {
		c.log = logger
		return nil
	}
}

func New(options ...Option) (*Client, error) {
	c := Client{
		httpClient: http.DefaultClient,
		log:        logrus.New(),
	}

	for _, option := range options {
		if err := option(&c); err != nil {
			return nil, err
		}
	}

	c.httpClient = newTransport(c.httpClient)

	return &c, nil
}
