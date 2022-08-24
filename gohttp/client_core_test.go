//go:build integration
// +build integration

package gohttp

import (
	"net/http"
	"testing"
)

func Test_GetRequestHeader(t *testing.T) {
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	if len(finalHeaders) != 3 {
		t.Error("se esperaban 3 headers")
	}
}

func Test_GetRequestBody(t *testing.T) {
	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)
		if err != nil {
			t.Error("no se esperaba error al enviar body nil")
		}

		if body != nil {
			t.Error("no se esperaba un body lleno al enviar body nil")

		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		rbody := []string{"uno", "dos"}
		body, err := client.getRequestBody("application/json", rbody)
		if err != nil {
			t.Error("no se esperaba error al serializar body")
		}

		if string(body) != `["uno","dos"]` {
			t.Error("body inválido")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		body, err := client.getRequestBody("application/xml", nil)
		if err != nil {
			t.Error("no se esperaba error al enviar body nil")
		}

		if body != nil {
			t.Error("no se esperaba un body lleno al enviar body nil")

		}
	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		body, err := client.getRequestBody("application/json", nil)
		if err != nil {
			t.Error("no se esperaba error al enviar body nil")
		}

		if body != nil {
			t.Error("no se esperaba un body lleno al enviar body nil")

		}
	})
}
