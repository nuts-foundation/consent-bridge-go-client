// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// ASymmetricKey defines model for ASymmetricKey.
type ASymmetricKey struct {
	Alg         *string    `json:"alg,omitempty"`
	CipherText  *string    `json:"cipherText,omitempty"`
	LegalEntity Identifier `json:"legalEntity"`
}

// ConsentId defines model for ConsentId.
type ConsentId struct {
	UUID       string  `json:"UUID"`
	ExternalId *string `json:"externalId,omitempty"`
}

// ConsentRecord defines model for ConsentRecord.
type ConsentRecord struct {
	AttachmentHash *string                     `json:"attachmentHash,omitempty"`
	CipherText     *string                     `json:"cipherText,omitempty"`
	Metadata       *Metadata                   `json:"metadata,omitempty"`
	Signatures     *[]PartyAttachmentSignature `json:"signatures,omitempty"`
}

// ConsentState defines model for ConsentState.
type ConsentState struct {
	ConsentId      ConsentId       `json:"consentId"`
	ConsentRecords []ConsentRecord `json:"consentRecords"`
}

// Domain defines model for Domain.
type Domain string

// FullConsentRequestState defines model for FullConsentRequestState.
type FullConsentRequestState struct {
	ConsentId      ConsentId       `json:"consentId"`
	ConsentRecords []ConsentRecord `json:"consentRecords"`
	LegalEntities  []Identifier    `json:"legalEntities"`
}

// Identifier defines model for Identifier.
type Identifier string

// Metadata defines model for Metadata.
type Metadata struct {
	Domain                 []Domain        `json:"domain"`
	OrganisationSecureKeys []ASymmetricKey `json:"organisationSecureKeys"`
	Period                 Period          `json:"period"`
	PreviousAttachmentHash *string         `json:"previousAttachmentHash,omitempty"`
	SecureKey              SymmetricKey    `json:"secureKey"`
}

// PartyAttachmentSignature defines model for PartyAttachmentSignature.
type PartyAttachmentSignature struct {
	Attachment  string           `json:"attachment"`
	LegalEntity Identifier       `json:"legalEntity"`
	Signature   SignatureWithKey `json:"signature"`
}

// Period defines model for Period.
type Period struct {
	ValidFrom time.Time  `json:"validFrom"`
	ValidTo   *time.Time `json:"validTo,omitempty"`
}

// SignatureWithKey defines model for SignatureWithKey.
type SignatureWithKey struct {
	Data      string `json:"data"`
	PublicKey string `json:"publicKey"`
}

// StateMachineId defines model for StateMachineId.
type StateMachineId string

// SymmetricKey defines model for SymmetricKey.
type SymmetricKey struct {
	Alg string `json:"alg"`
	Iv  string `json:"iv"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(req *http.Request, ctx context.Context) error

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// HTTP client with any customized settings, such as certificate chains.
	Client http.Client

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetAttachmentBySecureHash request
	GetAttachmentBySecureHash(ctx context.Context, secureHash string) (*http.Response, error)

	// GetConsentRequestStateById request
	GetConsentRequestStateById(ctx context.Context, uuid string) (*http.Response, error)
}

func (c *Client) GetAttachmentBySecureHash(ctx context.Context, secureHash string) (*http.Response, error) {
	req, err := NewGetAttachmentBySecureHashRequest(c.Server, secureHash)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetConsentRequestStateById(ctx context.Context, uuid string) (*http.Response, error) {
	req, err := NewGetConsentRequestStateByIdRequest(c.Server, uuid)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewGetAttachmentBySecureHashRequest generates requests for GetAttachmentBySecureHash
func NewGetAttachmentBySecureHashRequest(server string, secureHash string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "secureHash", secureHash)
	if err != nil {
		return nil, err
	}

	queryUrl := fmt.Sprintf("%s/api/attachment/%s", server, pathParam0)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetConsentRequestStateByIdRequest generates requests for GetConsentRequestStateById
func NewGetConsentRequestStateByIdRequest(server string, uuid string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "uuid", uuid)
	if err != nil {
		return nil, err
	}

	queryUrl := fmt.Sprintf("%s/api/consent_request_state/%s", server, pathParam0)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses returns a ClientWithResponses with a default Client:
func NewClientWithResponses(server string) *ClientWithResponses {
	return &ClientWithResponses{
		ClientInterface: &Client{
			Client: http.Client{},
			Server: server,
		},
	}
}

// NewClientWithResponsesAndRequestEditorFunc takes in a RequestEditorFn callback function and returns a ClientWithResponses with a default Client:
func NewClientWithResponsesAndRequestEditorFunc(server string, reqEditorFn RequestEditorFn) *ClientWithResponses {
	return &ClientWithResponses{
		ClientInterface: &Client{
			Client:        http.Client{},
			Server:        server,
			RequestEditor: reqEditorFn,
		},
	}
}

type getAttachmentBySecureHashResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r getAttachmentBySecureHashResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r getAttachmentBySecureHashResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type getConsentRequestStateByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FullConsentRequestState
}

// Status returns HTTPResponse.Status
func (r getConsentRequestStateByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r getConsentRequestStateByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAttachmentBySecureHashWithResponse request returning *GetAttachmentBySecureHashResponse
func (c *ClientWithResponses) GetAttachmentBySecureHashWithResponse(ctx context.Context, secureHash string) (*getAttachmentBySecureHashResponse, error) {
	rsp, err := c.GetAttachmentBySecureHash(ctx, secureHash)
	if err != nil {
		return nil, err
	}
	return ParsegetAttachmentBySecureHashResponse(rsp)
}

// GetConsentRequestStateByIdWithResponse request returning *GetConsentRequestStateByIdResponse
func (c *ClientWithResponses) GetConsentRequestStateByIdWithResponse(ctx context.Context, uuid string) (*getConsentRequestStateByIdResponse, error) {
	rsp, err := c.GetConsentRequestStateById(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return ParsegetConsentRequestStateByIdResponse(rsp)
}

// ParsegetAttachmentBySecureHashResponse parses an HTTP response from a GetAttachmentBySecureHashWithResponse call
func ParsegetAttachmentBySecureHashResponse(rsp *http.Response) (*getAttachmentBySecureHashResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &getAttachmentBySecureHashResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case rsp.StatusCode == 200:
	// Content-type (application/octet-stream) unsupported
	case rsp.StatusCode == 404:
		break // No content-type
	}

	return response, nil
}

// ParsegetConsentRequestStateByIdResponse parses an HTTP response from a GetConsentRequestStateByIdWithResponse call
func ParsegetConsentRequestStateByIdResponse(rsp *http.Response) (*getConsentRequestStateByIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &getConsentRequestStateByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		response.JSON200 = &FullConsentRequestState{}
		if err := json.Unmarshal(bodyBytes, response.JSON200); err != nil {
			return nil, err
		}
	case rsp.StatusCode == 404:
		break // No content-type
	}

	return response, nil
}

