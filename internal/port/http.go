package port

import (
	"errors"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5/middleware"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
	"net/http"
	"time"
)

//go:generate oapi-codegen --config=../../openapi_server.yaml ../../api/openapi/user-service.yaml
//go:generate oapi-codegen --config=../../openapi_types.yaml ../../api/openapi/user-service.yaml

const (
	defaultHTTPReadTimeout       = 5 * time.Second
	defaultHTTPReadHeaderTimeout = 5 * time.Second
	defaultHTTPWriteTimeout      = 5 * time.Second
)

func NewServer(address string, service StrictServerInterface) (*http.Server, error) {
	if address == "" {
		return nil, ErrInvalidHTTPPort
	}

	spec, err := GetSwagger()
	if err != nil {
		return nil, err
	}

	spec.Servers = nil

	handler := NewStrictHandler(service, nil)

	options := StdHTTPServerOptions{
		BaseRouter: http.NewServeMux(),
		Middlewares: []MiddlewareFunc{
			nethttpmiddleware.OapiRequestValidatorWithOptions(spec, &nethttpmiddleware.Options{
				SilenceServersWarning: true,
				ErrorHandler:          sendErrorResponse,
				MultiErrorHandler:     handleValidationError,
				Options: openapi3filter.Options{
					MultiError:         true,
					AuthenticationFunc: authenticate,
				},
			}),
			requestLoggerMiddleware,
			middleware.RealIP,
			middleware.NoCache,
			middleware.SetHeader("Content-Type", "application/json"),
			middleware.SetHeader("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload"),
			middleware.SetHeader("X-Frame-Options", "DENY"),
			middleware.SetHeader("X-Content-Type-Options", "nosniff"),
			middleware.SetHeader("Access-Control-Allow-Origin", "*"),
			middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS"),
			middleware.SetHeader("Access-Control-Allow-Headers", "content-type, authorization, request-id, traceparent, request-context"),
			middleware.Recoverer,
		},
	}

	server := http.Server{
		Addr:              address,
		Handler:           HandlerWithOptions(handler, options),
		ReadTimeout:       defaultHTTPReadTimeout,
		ReadHeaderTimeout: defaultHTTPReadHeaderTimeout,
		WriteTimeout:      defaultHTTPWriteTimeout,
	}

	return &server, nil
}

func handleValidationError(me openapi3.MultiError) (int, error) {
	securityError := &openapi3filter.SecurityRequirementsError{}

	if !me.As(&securityError) {
		return http.StatusBadRequest, me
	}

	for i := range securityError.Errors {
		if errors.Is(securityError.Errors[i], ErrNoAuthHeader) {
			return http.StatusUnauthorized, me
		}

		if errors.Is(securityError.Errors[i], ErrGetBearerToken) {
			return http.StatusForbidden, me
		}
	}

	return http.StatusBadRequest, me
}
