package gohttp

import (
	"crypto/tls"
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
	client             *http.Client
	userAgent          string
	tlsconfig          *tls.Config
}

type ClientBuilder interface {
	SetCommonHeader(header http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	SetHttpClient(c *http.Client) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder
	SetTLSClientConfig(tlsConfig *tls.Config) ClientBuilder
	Build() Client
}

// NewBuilder genera un cliente http
func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

// Build genera un cliente http
func (c *clientBuilder) Build() Client {

	client := httpClient{
		builder: c,
	}

	return &client
}

// SetCommonHeader inicializa los headers comunes para cada petición el cliente
func (c *clientBuilder) SetCommonHeader(header http.Header) ClientBuilder {
	c.headers = header
	return c
}

// SetConnectionTimeout asigna el tiempo de espera para conexiones del cliente
func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

// SetResponseTimeout asigna el tiempo de espera para respuestas a peticiones realizadas por el cliente
func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

// SetMaxIdleConnections asigna el número máximo de conexiones vacías o en espera
func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnections = i
	return c
}

// SetCommonHeader inicializa los headers comunes para cada petición el cliente
func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

// SetUserAgent asigna el agente para el cliente
func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	c.userAgent = userAgent
	return c
}

// SetTLSClientConfig asigna la configuración de seguridad para el cliente
func (c *clientBuilder) SetTLSClientConfig(tlsConfig *tls.Config) ClientBuilder {
	c.tlsconfig = tlsConfig
	return c
}

// DisableTimeouts deshabilita los timeouts
func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}
