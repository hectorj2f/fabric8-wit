// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "tenant": tenant Resource Client
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
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// SetupTenantPath computes a request path to the setup action of tenant.
func SetupTenantPath() string {

	return fmt.Sprintf("/api/tenant")
}

// Initialize new tenant environment.
func (c *Client) SetupTenant(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSetupTenantRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSetupTenantRequest create the request corresponding to the setup action endpoint of the tenant resource.
func (c *Client) NewSetupTenantRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ShowTenantPath computes a request path to the show action of tenant.
func ShowTenantPath() string {

	return fmt.Sprintf("/api/tenant")
}

// Initialize new tenant environment.
func (c *Client) ShowTenant(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowTenantRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowTenantRequest create the request corresponding to the show action endpoint of the tenant resource.
func (c *Client) NewShowTenantRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// UpdateTenantPath computes a request path to the update action of tenant.
func UpdateTenantPath() string {

	return fmt.Sprintf("/api/tenant")
}

// Initialize new tenant environment.
func (c *Client) UpdateTenant(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewUpdateTenantRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateTenantRequest create the request corresponding to the update action endpoint of the tenant resource.
func (c *Client) NewUpdateTenantRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
