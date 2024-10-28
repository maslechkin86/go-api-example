package port

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3filter"
	"log"
	"net/http"
	"slices"
	"strings"
)

const (
	bearerPrefix string = "Bearer "
)

type (
	responseData struct {
		status        int
		contentLength int64
	}

	responseDataWriter struct {
		http.ResponseWriter
		responseData responseData
	}
)

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	o := ErrorResponse{Message: message, Code: statusCode}

	buf, err := json.Marshal(&o)
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf("{\"code\": %d, \"message\": \"%s\"}", statusCode, err.Error())))
	} else {
		_, _ = w.Write(buf)
	}
}

func requestLoggerMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		writer := responseDataWriter{ResponseWriter: w}
		h.ServeHTTP(&writer, r)

		ignoreRouts := []string{}
		if !slices.Contains(ignoreRouts, r.URL.String()) {
			log.Printf("processed request %s: %s [%d]", r.Method, r.URL.String(), writer.responseData.status)
		}
	}

	return http.HandlerFunc(fn)
}

func authenticate(_ context.Context, input *openapi3filter.AuthenticationInput) error {
	_, err := getBearerTokenFromRequest(input.RequestValidationInput.Request)
	return err
}

func getBearerTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")

	if len(header) == 0 {
		return "", ErrNoAuthHeader
	}

	if !strings.HasPrefix(header, bearerPrefix) {
		return "", ErrGetBearerToken
	}

	return strings.TrimPrefix(header, bearerPrefix), nil
}
