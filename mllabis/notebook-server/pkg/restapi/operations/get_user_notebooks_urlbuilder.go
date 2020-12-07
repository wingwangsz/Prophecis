// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetUserNotebooksURL generates an URL for the get user notebooks operation
type GetUserNotebooksURL struct {
	User string

	ClusterName *string
	Page        *string
	Size        *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUserNotebooksURL) WithBasePath(bp string) *GetUserNotebooksURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUserNotebooksURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetUserNotebooksURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/aide/v1/user/{user}/notebooks"

	user := o.User
	if user != "" {
		_path = strings.Replace(_path, "{user}", user, -1)
	} else {
		return nil, errors.New("user is required on GetUserNotebooksURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var clusterName string
	if o.ClusterName != nil {
		clusterName = *o.ClusterName
	}
	if clusterName != "" {
		qs.Set("clusterName", clusterName)
	}

	var page string
	if o.Page != nil {
		page = *o.Page
	}
	if page != "" {
		qs.Set("page", page)
	}

	var size string
	if o.Size != nil {
		size = *o.Size
	}
	if size != "" {
		qs.Set("size", size)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetUserNotebooksURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetUserNotebooksURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetUserNotebooksURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetUserNotebooksURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetUserNotebooksURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetUserNotebooksURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
