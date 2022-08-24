package gohttp

import (
	"net/http"
	"sync"

	"github.com/RaulGarciaMz/go-httpclient/core"
)

// httpClient implementa la interface del cliente Http
type httpClient struct {
	builder *clientBuilder

	client     *http.Client
	clientOnce sync.Once
}

// Client define las operaciones del cliente Http
type Client interface {
	Get(url string, headers ...http.Header) (*core.Response, error)
	GetWithBody(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Delete(url string, headers ...http.Header) (*core.Response, error)
	Options(url string, headers ...http.Header) (*core.Response, error)
}

// Get implementa las peticiones Get del cliente Http
func (c *httpClient) Get(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), nil)
}

// GetWithBody implementa las peticiones Post del cliente Http
func (c *httpClient) GetWithBody(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), body)
}

// Post implementa las peticiones Post del cliente Http
func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

// Put implementa las peticiones Put del cliente Http
func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), body)
}

// Patch implementa las peticiones Patch del cliente Http
func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), body)
}

// Delete implementa las peticiones Delete del cliente Http
func (c *httpClient) Delete(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}

// Options implementa las peticiones Options del cliente Http
func (c *httpClient) Options(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodOptions, url, getHeaders(headers...), nil)
}
