//go:build go1.22

// Package port provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package port

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// The health check endpoint
	// (GET /v1/healthz)
	HealthzCheck(w http.ResponseWriter, r *http.Request)
	// Retrieve users list
	// (GET /v1/users)
	GetUserList(w http.ResponseWriter, r *http.Request, params GetUserListParams)
	// Create a new user
	// (POST /v1/users)
	CreateUser(w http.ResponseWriter, r *http.Request)
	// Retrieve user details
	// (GET /v1/users/{id})
	GetUser(w http.ResponseWriter, r *http.Request, id int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// HealthzCheck operation middleware
func (siw *ServerInterfaceWrapper) HealthzCheck(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.HealthzCheck(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUserList operation middleware
func (siw *ServerInterfaceWrapper) GetUserList(w http.ResponseWriter, r *http.Request) {

	var err error

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserListParams

	// ------------- Required query parameter "limit" -------------

	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Required query parameter "skip" -------------

	if paramValue := r.URL.Query().Get("skip"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "skip"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "skip", r.URL.Query(), &params.Skip)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "skip", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserList(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUser(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/v1/healthz", wrapper.HealthzCheck)
	m.HandleFunc("GET "+options.BaseURL+"/v1/users", wrapper.GetUserList)
	m.HandleFunc("POST "+options.BaseURL+"/v1/users", wrapper.CreateUser)
	m.HandleFunc("GET "+options.BaseURL+"/v1/users/{id}", wrapper.GetUser)

	return m
}

type HealthzCheckRequestObject struct {
}

type HealthzCheckResponseObject interface {
	VisitHealthzCheckResponse(w http.ResponseWriter) error
}

type HealthzCheck200JSONResponse HealthCheckResponse

func (response HealthzCheck200JSONResponse) VisitHealthzCheckResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUserListRequestObject struct {
	Params GetUserListParams
}

type GetUserListResponseObject interface {
	VisitGetUserListResponse(w http.ResponseWriter) error
}

type GetUserList200JSONResponse UserListResponse

func (response GetUserList200JSONResponse) VisitGetUserListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUserList400JSONResponse ErrorResponse

func (response GetUserList400JSONResponse) VisitGetUserListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetUserList401JSONResponse ErrorResponse

func (response GetUserList401JSONResponse) VisitGetUserListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetUserList403JSONResponse ErrorResponse

func (response GetUserList403JSONResponse) VisitGetUserListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type GetUserList404JSONResponse ErrorResponse

func (response GetUserList404JSONResponse) VisitGetUserListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetUserList409JSONResponse ErrorResponse

func (response GetUserList409JSONResponse) VisitGetUserListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type GetUserList500JSONResponse ErrorResponse

func (response GetUserList500JSONResponse) VisitGetUserListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetUserList502JSONResponse ErrorResponse

func (response GetUserList502JSONResponse) VisitGetUserListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(502)

	return json.NewEncoder(w).Encode(response)
}

type CreateUserRequestObject struct {
	Body *CreateUserJSONRequestBody
}

type CreateUserResponseObject interface {
	VisitCreateUserResponse(w http.ResponseWriter) error
}

type CreateUser201JSONResponse User

func (response CreateUser201JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser400JSONResponse ErrorResponse

func (response CreateUser400JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser401JSONResponse ErrorResponse

func (response CreateUser401JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser403JSONResponse ErrorResponse

func (response CreateUser403JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser409JSONResponse ErrorResponse

func (response CreateUser409JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser500JSONResponse ErrorResponse

func (response CreateUser500JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser502JSONResponse ErrorResponse

func (response CreateUser502JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(502)

	return json.NewEncoder(w).Encode(response)
}

type GetUserRequestObject struct {
	ID int `json:"id"`
}

type GetUserResponseObject interface {
	VisitGetUserResponse(w http.ResponseWriter) error
}

type GetUser200JSONResponse User

func (response GetUser200JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUser400JSONResponse ErrorResponse

func (response GetUser400JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetUser401JSONResponse ErrorResponse

func (response GetUser401JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetUser403JSONResponse ErrorResponse

func (response GetUser403JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type GetUser404JSONResponse ErrorResponse

func (response GetUser404JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetUser409JSONResponse ErrorResponse

func (response GetUser409JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type GetUser500JSONResponse ErrorResponse

func (response GetUser500JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetUser502JSONResponse ErrorResponse

func (response GetUser502JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(502)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// The health check endpoint
	// (GET /v1/healthz)
	HealthzCheck(ctx context.Context, request HealthzCheckRequestObject) (HealthzCheckResponseObject, error)
	// Retrieve users list
	// (GET /v1/users)
	GetUserList(ctx context.Context, request GetUserListRequestObject) (GetUserListResponseObject, error)
	// Create a new user
	// (POST /v1/users)
	CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error)
	// Retrieve user details
	// (GET /v1/users/{id})
	GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// HealthzCheck operation middleware
func (sh *strictHandler) HealthzCheck(w http.ResponseWriter, r *http.Request) {
	var request HealthzCheckRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.HealthzCheck(ctx, request.(HealthzCheckRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "HealthzCheck")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(HealthzCheckResponseObject); ok {
		if err := validResponse.VisitHealthzCheckResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetUserList operation middleware
func (sh *strictHandler) GetUserList(w http.ResponseWriter, r *http.Request, params GetUserListParams) {
	var request GetUserListRequestObject

	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetUserList(ctx, request.(GetUserListRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUserList")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetUserListResponseObject); ok {
		if err := validResponse.VisitGetUserListResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateUser operation middleware
func (sh *strictHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequestObject

	var body CreateUserJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CreateUser(ctx, request.(CreateUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateUser")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CreateUserResponseObject); ok {
		if err := validResponse.VisitCreateUserResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetUser operation middleware
func (sh *strictHandler) GetUser(w http.ResponseWriter, r *http.Request, id int) {
	var request GetUserRequestObject

	request.ID = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetUser(ctx, request.(GetUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUser")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetUserResponseObject); ok {
		if err := validResponse.VisitGetUserResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYXW/bNhT9KwS3hw1Q/dFmRaa3NG1aD91WJCn2kPiBFq+t20ikQl45cwL994GkZDmx",
	"VLdbahRDnmxR1L3nfh0e6Y4nOi+0AkWWx3fcJinkwv99Y4w2p2ALrSy4hcLoAgwh+NuJln5Vgk0MFoRa",
	"8Zifp8DenZ9/YJYElZa5XQMecVoVwGOOimABhlcRz8FasegwccTW10tg4FCwevOGJUsG1YJXVcQNXJdo",
	"QPL4IqBqjU/X+/XsEyTkHL8DkVF6nEJy1R9dgN8dX+oNsMRZaOL8CQaLQcQuub665O7X477kP3dAjjhh",
	"DpZEXnQ7SEpjQBFbb2Oo2OTsT3b4cjRmc21yQc5u+MdjLgXBM7d7Z37quDYhdOXoowWznRSU3XhLC4ah",
	"BEU4RzAd9Y74388W+pkSuVudvHY+wkWvPXd7d71R8tpSXxjHBgTBKVyXYGk7pkdB8VkA79FSf6MhQd7T",
	"ZxlaYnrugVgHYr33RwNzHvMfhu3wDuvJHfrSVWsswhixctcZ5kh9jnL0nrwDdpNikjIDVBoFkhVgmAnp",
	"655ke4U9jSxyXaoNy6SZ29xthjSJbNKfDn8/GOoy8KAkG9aa2GukTR636+VigaQ0SKszl81QoRkIA+ao",
	"pLS9Omkm77e/zp1dv5vH9d0WXUpU8MoZRjXXXXGhZUcfJkxkmb6xbrYZpcByocQCclD3G4CQMmfW1Zid",
	"gVli4kZ+CcYGe+PBaDByydQFKFEgj/mLwWgw5hEvBKU+nuFyPAwUdusuF9DRFW+ULDQ6CtI1z1FLfDXl",
	"6blftAGGg+f6WjgLE8njmmdvPdFyV5wwAh7D89EonCGKQHn/oigyTPzTw0/WgWgOo10N38XnPuXbLVRj",
	"ZWjrWFaDUPUyz4VZdfE71Jnw+1zqfDG+KHEGyCAswadJZBkTS4GZmGXQVvR+yt4CNYzhK2ZEDuTdXfzn",
	"uUX31HUJZtUwZryei3ZwyJQQbSRewlyUGfH4+SjiOSrMy5zHo6759QzfLr884FX0laTQhbGe2S+A+G8Q",
	"Tr9hY26Rf0dX+kn2RN90i2S2TBKwdl5mmevPiB88Iqj7qq4D0UQtRYaSoSpKYlKQCBDG+4PwUYmSUm3w",
	"FiQTPhkBw4v9YTjRZoZSgmIi8UsewMH+APyhiZ3oUsng+df9eT7Wap5hEDC/7Lf1CIwSmT/awDD/QEDx",
	"fH8oXgnJ3gqCG6ebNiSBJ+FNMXAxdfTRnh2nDd97dvdT7c5ssXD8HSb99/W5zqdVxAttdx2+Xr0ywRTc",
	"eLvbh0YQuF72BZoES6+0XD0qj91X0T3HqwQSmK2VgRfQpNkM6jDkYIvIqy3+HT8q7l7OrRHdI9snrv2e",
	"uPaJ8b57xjt+yE6f5btNAT28Q1l9vYr2lNLQzGzFkCybvO4V0vwbq7tedmkgPom6J1H3JOr+N6Kumetd",
	"sm7Xi3v7rdRpJbEWdv79txCUtq+//gNn/8tvz0vuxofWaQjX5b4LznudiIxHvDRZ/b0qHg4zt5hqS/Hh",
	"6HDk35LreB8+HiJvv2fUjh9mpZpW/wQAAP///Ve85moYAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
