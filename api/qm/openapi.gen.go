// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateIssueRequest defines model for CreateIssueRequest.
type CreateIssueRequest struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Issue defines model for Issue.
type Issue struct {
	// Id A 24 character hexadecimal unique material Identifier.
	Id        *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Action    *string             `bson:"action,omitempty" json:"action,omitempty"`
	ChangeReq *bool               `bson:"change_req,omitempty" json:"change_req,omitempty"`

	// Description indicates the description about the issue
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	Impact      *string `bson:"impact,omitempty" json:"impact,omitempty"`

	// Name indicates the name of the issue
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
}

// IssueBase defines model for IssueBase.
type IssueBase struct {
	Action    *string `bson:"action,omitempty" json:"action,omitempty"`
	ChangeReq *bool   `bson:"change_req,omitempty" json:"change_req,omitempty"`

	// Description indicates the description about the issue
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	Impact      *string `bson:"impact,omitempty" json:"impact,omitempty"`

	// Name indicates the name of the issue
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
}

// UpdateIssueRequest defines model for UpdateIssueRequest.
type UpdateIssueRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// PageNumber defines model for PageNumber.
type PageNumber = int

// PageSize defines model for PageSize.
type PageSize = int

// GetAllIssuesParams defines parameters for GetAllIssues.
type GetAllIssuesParams struct {
	// Page Page number
	Page *PageNumber `form:"page,omitempty" json:"page,omitempty"`

	// PageSize Number of items per page
	PageSize *PageSize `form:"pageSize,omitempty" json:"pageSize,omitempty"`
}

// CreateIssueJSONRequestBody defines body for CreateIssue for application/json ContentType.
type CreateIssueJSONRequestBody = CreateIssueRequest

// UpdateIssueByIdJSONRequestBody defines body for UpdateIssueById for application/json ContentType.
type UpdateIssueByIdJSONRequestBody = UpdateIssueRequest

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetAllIssues request
	GetAllIssues(ctx context.Context, params *GetAllIssuesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateIssueWithBody request with any body
	CreateIssueWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateIssue(ctx context.Context, body CreateIssueJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteIssueById request
	DeleteIssueById(ctx context.Context, issueId openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetIssueById request
	GetIssueById(ctx context.Context, issueId openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateIssueByIdWithBody request with any body
	UpdateIssueByIdWithBody(ctx context.Context, issueId openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateIssueById(ctx context.Context, issueId openapi_types.UUID, body UpdateIssueByIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAllIssues(ctx context.Context, params *GetAllIssuesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAllIssuesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateIssueWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateIssueRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateIssue(ctx context.Context, body CreateIssueJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateIssueRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteIssueById(ctx context.Context, issueId openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteIssueByIdRequest(c.Server, issueId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetIssueById(ctx context.Context, issueId openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetIssueByIdRequest(c.Server, issueId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateIssueByIdWithBody(ctx context.Context, issueId openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateIssueByIdRequestWithBody(c.Server, issueId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateIssueById(ctx context.Context, issueId openapi_types.UUID, body UpdateIssueByIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateIssueByIdRequest(c.Server, issueId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAllIssuesRequest generates requests for GetAllIssues
func NewGetAllIssuesRequest(server string, params *GetAllIssuesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/itr")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Page != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.PageSize != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "pageSize", runtime.ParamLocationQuery, *params.PageSize); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateIssueRequest calls the generic CreateIssue builder with application/json body
func NewCreateIssueRequest(server string, body CreateIssueJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateIssueRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateIssueRequestWithBody generates requests for CreateIssue with any type of body
func NewCreateIssueRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/itr")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteIssueByIdRequest generates requests for DeleteIssueById
func NewDeleteIssueByIdRequest(server string, issueId openapi_types.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "issueId", runtime.ParamLocationPath, issueId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/itr/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetIssueByIdRequest generates requests for GetIssueById
func NewGetIssueByIdRequest(server string, issueId openapi_types.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "issueId", runtime.ParamLocationPath, issueId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/itr/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateIssueByIdRequest calls the generic UpdateIssueById builder with application/json body
func NewUpdateIssueByIdRequest(server string, issueId openapi_types.UUID, body UpdateIssueByIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateIssueByIdRequestWithBody(server, issueId, "application/json", bodyReader)
}

// NewUpdateIssueByIdRequestWithBody generates requests for UpdateIssueById with any type of body
func NewUpdateIssueByIdRequestWithBody(server string, issueId openapi_types.UUID, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "issueId", runtime.ParamLocationPath, issueId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/itr/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetAllIssuesWithResponse request
	GetAllIssuesWithResponse(ctx context.Context, params *GetAllIssuesParams, reqEditors ...RequestEditorFn) (*GetAllIssuesResponse, error)

	// CreateIssueWithBodyWithResponse request with any body
	CreateIssueWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateIssueResponse, error)

	CreateIssueWithResponse(ctx context.Context, body CreateIssueJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateIssueResponse, error)

	// DeleteIssueByIdWithResponse request
	DeleteIssueByIdWithResponse(ctx context.Context, issueId openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteIssueByIdResponse, error)

	// GetIssueByIdWithResponse request
	GetIssueByIdWithResponse(ctx context.Context, issueId openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetIssueByIdResponse, error)

	// UpdateIssueByIdWithBodyWithResponse request with any body
	UpdateIssueByIdWithBodyWithResponse(ctx context.Context, issueId openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateIssueByIdResponse, error)

	UpdateIssueByIdWithResponse(ctx context.Context, issueId openapi_types.UUID, body UpdateIssueByIdJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateIssueByIdResponse, error)
}

type GetAllIssuesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Issue
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GetAllIssuesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllIssuesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateIssueResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *Issue
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r CreateIssueResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateIssueResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteIssueByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r DeleteIssueByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteIssueByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetIssueByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Issue
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GetIssueByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetIssueByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateIssueByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Issue
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r UpdateIssueByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateIssueByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAllIssuesWithResponse request returning *GetAllIssuesResponse
func (c *ClientWithResponses) GetAllIssuesWithResponse(ctx context.Context, params *GetAllIssuesParams, reqEditors ...RequestEditorFn) (*GetAllIssuesResponse, error) {
	rsp, err := c.GetAllIssues(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAllIssuesResponse(rsp)
}

// CreateIssueWithBodyWithResponse request with arbitrary body returning *CreateIssueResponse
func (c *ClientWithResponses) CreateIssueWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateIssueResponse, error) {
	rsp, err := c.CreateIssueWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateIssueResponse(rsp)
}

func (c *ClientWithResponses) CreateIssueWithResponse(ctx context.Context, body CreateIssueJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateIssueResponse, error) {
	rsp, err := c.CreateIssue(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateIssueResponse(rsp)
}

// DeleteIssueByIdWithResponse request returning *DeleteIssueByIdResponse
func (c *ClientWithResponses) DeleteIssueByIdWithResponse(ctx context.Context, issueId openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteIssueByIdResponse, error) {
	rsp, err := c.DeleteIssueById(ctx, issueId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteIssueByIdResponse(rsp)
}

// GetIssueByIdWithResponse request returning *GetIssueByIdResponse
func (c *ClientWithResponses) GetIssueByIdWithResponse(ctx context.Context, issueId openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetIssueByIdResponse, error) {
	rsp, err := c.GetIssueById(ctx, issueId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetIssueByIdResponse(rsp)
}

// UpdateIssueByIdWithBodyWithResponse request with arbitrary body returning *UpdateIssueByIdResponse
func (c *ClientWithResponses) UpdateIssueByIdWithBodyWithResponse(ctx context.Context, issueId openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateIssueByIdResponse, error) {
	rsp, err := c.UpdateIssueByIdWithBody(ctx, issueId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateIssueByIdResponse(rsp)
}

func (c *ClientWithResponses) UpdateIssueByIdWithResponse(ctx context.Context, issueId openapi_types.UUID, body UpdateIssueByIdJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateIssueByIdResponse, error) {
	rsp, err := c.UpdateIssueById(ctx, issueId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateIssueByIdResponse(rsp)
}

// ParseGetAllIssuesResponse parses an HTTP response from a GetAllIssuesWithResponse call
func ParseGetAllIssuesResponse(rsp *http.Response) (*GetAllIssuesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAllIssuesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Issue
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseCreateIssueResponse parses an HTTP response from a CreateIssueWithResponse call
func ParseCreateIssueResponse(rsp *http.Response) (*CreateIssueResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateIssueResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest Issue
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDeleteIssueByIdResponse parses an HTTP response from a DeleteIssueByIdWithResponse call
func ParseDeleteIssueByIdResponse(rsp *http.Response) (*DeleteIssueByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteIssueByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetIssueByIdResponse parses an HTTP response from a GetIssueByIdWithResponse call
func ParseGetIssueByIdResponse(rsp *http.Response) (*GetIssueByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetIssueByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Issue
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseUpdateIssueByIdResponse parses an HTTP response from a UpdateIssueByIdWithResponse call
func ParseUpdateIssueByIdResponse(rsp *http.Response) (*UpdateIssueByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateIssueByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Issue
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
