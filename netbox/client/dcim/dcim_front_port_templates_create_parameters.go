// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/mikeburke14/go-netbox/netbox/models"
)

// NewDcimFrontPortTemplatesCreateParams creates a new DcimFrontPortTemplatesCreateParams object
// with the default values initialized.
func NewDcimFrontPortTemplatesCreateParams() *DcimFrontPortTemplatesCreateParams {
	var ()
	return &DcimFrontPortTemplatesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesCreateParamsWithTimeout creates a new DcimFrontPortTemplatesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimFrontPortTemplatesCreateParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesCreateParams {
	var ()
	return &DcimFrontPortTemplatesCreateParams{

		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesCreateParamsWithContext creates a new DcimFrontPortTemplatesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimFrontPortTemplatesCreateParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesCreateParams {
	var ()
	return &DcimFrontPortTemplatesCreateParams{

		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesCreateParamsWithHTTPClient creates a new DcimFrontPortTemplatesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimFrontPortTemplatesCreateParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesCreateParams {
	var ()
	return &DcimFrontPortTemplatesCreateParams{
		HTTPClient: client,
	}
}

/*DcimFrontPortTemplatesCreateParams contains all the parameters to send to the API endpoint
for the dcim front port templates create operation typically these are written to a http.Request
*/
type DcimFrontPortTemplatesCreateParams struct {

	/*Data*/
	Data *models.WritableFrontPortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) WithData(data *models.WritableFrontPortTemplate) *DcimFrontPortTemplatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) SetData(data *models.WritableFrontPortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
