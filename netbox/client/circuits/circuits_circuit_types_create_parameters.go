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

package circuits

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

// NewCircuitsCircuitTypesCreateParams creates a new CircuitsCircuitTypesCreateParams object
// with the default values initialized.
func NewCircuitsCircuitTypesCreateParams() *CircuitsCircuitTypesCreateParams {
	var ()
	return &CircuitsCircuitTypesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesCreateParamsWithTimeout creates a new CircuitsCircuitTypesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitTypesCreateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesCreateParams {
	var ()
	return &CircuitsCircuitTypesCreateParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesCreateParamsWithContext creates a new CircuitsCircuitTypesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitTypesCreateParamsWithContext(ctx context.Context) *CircuitsCircuitTypesCreateParams {
	var ()
	return &CircuitsCircuitTypesCreateParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitTypesCreateParamsWithHTTPClient creates a new CircuitsCircuitTypesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitTypesCreateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesCreateParams {
	var ()
	return &CircuitsCircuitTypesCreateParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitTypesCreateParams contains all the parameters to send to the API endpoint
for the circuits circuit types create operation typically these are written to a http.Request
*/
type CircuitsCircuitTypesCreateParams struct {

	/*Data*/
	Data *models.CircuitType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuit types create params
func (o *CircuitsCircuitTypesCreateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types create params
func (o *CircuitsCircuitTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types create params
func (o *CircuitsCircuitTypesCreateParams) WithContext(ctx context.Context) *CircuitsCircuitTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types create params
func (o *CircuitsCircuitTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types create params
func (o *CircuitsCircuitTypesCreateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types create params
func (o *CircuitsCircuitTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit types create params
func (o *CircuitsCircuitTypesCreateParams) WithData(data *models.CircuitType) *CircuitsCircuitTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit types create params
func (o *CircuitsCircuitTypesCreateParams) SetData(data *models.CircuitType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
