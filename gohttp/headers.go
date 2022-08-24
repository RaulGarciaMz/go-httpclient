package gohttp

import (
	"net/http"

	"github.com/RaulGarciaMz/go-httpclient/gomime"
)

func getHeaders(headers ...http.Header) http.Header {

	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

// getRequestHeaders genera el encabezado de la petición usando los encabezados comunes y los personalizados
func (c *httpClient) getRequestHeaders(requestHeader http.Header) http.Header {
	result := make(http.Header)

	//Headers comunes a todas las peticiones
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	//Headers personalizados en la petición
	for header, value := range requestHeader {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	//Asigna el User-Agent header para la petición
	if c.builder.userAgent != "" {
		if result.Get(gomime.HeaderUserAgent) != "" {
			return result
		}
		result.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}
	return result
}
