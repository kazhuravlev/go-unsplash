package unsplash

import "net/http"

const apiVersion = "v1"

type Transport struct {
	base http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept-Version", apiVersion)

	return t.base.RoundTrip(req)
}

func newTransport(c *http.Client) *http.Client {
	return &http.Client{
		Transport: &Transport{
			base: c.Transport,
		},
	}
}
