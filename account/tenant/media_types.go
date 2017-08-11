// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "tenant": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-tenant/design
// --notool=true
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-wit/account
// --pkg=tenant
// --version=v1.2.0

package tenant

import (
	"github.com/goadesign/goa"
	"net/http"
)

// JSONAPIErrors media type (default view)
//
// Identifier: application/vnd.jsonapierrors+json; view=default
type JSONAPIErrors struct {
	Errors []*JSONAPIError `form:"errors" json:"errors" xml:"errors"`
}

// Validate validates the JSONAPIErrors media type instance.
func (mt *JSONAPIErrors) Validate() (err error) {
	if mt.Errors == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "errors"))
	}
	for _, e := range mt.Errors {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeJSONAPIErrors decodes the JSONAPIErrors instance encoded in resp body.
func (c *Client) DecodeJSONAPIErrors(resp *http.Response) (*JSONAPIErrors, error) {
	var decoded JSONAPIErrors
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The status of the current running instance (default view)
//
// Identifier: application/vnd.status+json; view=default
type Status struct {
	// The time when built
	BuildTime string `form:"buildTime" json:"buildTime" xml:"buildTime"`
	// Commit SHA this build is based on
	Commit string `form:"commit" json:"commit" xml:"commit"`
	// The error if any
	Error *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
	// The time when started
	StartTime string `form:"startTime" json:"startTime" xml:"startTime"`
}

// Validate validates the Status media type instance.
func (mt *Status) Validate() (err error) {
	if mt.Commit == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "commit"))
	}
	if mt.BuildTime == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "buildTime"))
	}
	if mt.StartTime == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "startTime"))
	}
	return
}

// DecodeStatus decodes the Status instance encoded in resp body.
func (c *Client) DecodeStatus(resp *http.Response) (*Status, error) {
	var decoded Status
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Holds a single Tenant (default view)
//
// Identifier: application/vnd.tenant+json; view=default
type TenantSingle struct {
	Data *Tenant `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// Validate validates the TenantSingle media type instance.
func (mt *TenantSingle) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if mt.Data != nil {
		if err2 := mt.Data.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeTenantSingle decodes the TenantSingle instance encoded in resp body.
func (c *Client) DecodeTenantSingle(resp *http.Response) (*TenantSingle, error) {
	var decoded TenantSingle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
