package gohttp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/RaulGarciaMz/go-httpclient/core"
	"github.com/RaulGarciaMz/go-httpclient/gohttp_mock"
	"github.com/RaulGarciaMz/go-httpclient/gomime"
)

const (
	defaultIdleConnections   = 5
	defaultResponseTimeout   = 5 * time.Second
	defaultConnectionTimeout = 50 * time.Millisecond
)

// do Realiza el llamado genérico a un cliente http
func (c *httpClient) do(method, url string, headers http.Header, body interface{}) (*core.Response, error) {

	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("no fue capaz de crear un request")
	}

	request.Header = fullHeaders

	client := c.getHttpClient()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	finalResponse := core.Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
		Body:       responseBody,
	}

	return &finalResponse, nil
}

// getRequestBody genera el body de la petición acorde al valor del encabezado de Content-Type
func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case gomime.ContentTypeJson:
		return json.Marshal(body)
	case gomime.ContentTypeXml:
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

// getHttpClient obtiene un cliente Http con configuraciones asignadas
func (c *httpClient) getHttpClient() core.HttpClient {

	if gohttp_mock.MockupServer.IsEnabled() {
		return gohttp_mock.MockupServer.GetMockedClient()
	}

	c.clientOnce.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}

		c.client = &http.Client{
			Timeout: c.getResponseTimeout() + c.getConnectionTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
				TLSClientConfig: c.getTlsConfig(),
			},
		}
	})

	return c.client
}

// getMaxIdleConnections obtiene el valor para MaxIdleConnection, si no está definido toma el valor de default
func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}

	return defaultIdleConnections
}

// getResponseTimeout obtiene el valor para el RequestTimeout, si no está definido toma el valor de default
func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}

	if c.builder.disableTimeouts {
		return 0
	}

	return defaultResponseTimeout
}

// getConnectionTimeout obtiene el valor para el ConnectionTimeout, si no está definido toma el valor de default
func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}

	if c.builder.disableTimeouts {
		return 0
	}

	return defaultConnectionTimeout
}

// getTlsConfig obtiene el valor para el TlsClientConfig, si no está definido toma el valor de default
func (c *httpClient) getTlsConfig() *tls.Config {
	if c.builder.tlsconfig != nil {
		return c.builder.tlsconfig
	}

	return nil
}
