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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mikeburke14/go-netbox/netbox/models"
)

// ExtrasExportTemplatesPartialUpdateReader is a Reader for the ExtrasExportTemplatesPartialUpdate structure.
type ExtrasExportTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasExportTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasExportTemplatesPartialUpdateOK creates a ExtrasExportTemplatesPartialUpdateOK with default headers values
func NewExtrasExportTemplatesPartialUpdateOK() *ExtrasExportTemplatesPartialUpdateOK {
	return &ExtrasExportTemplatesPartialUpdateOK{}
}

/*ExtrasExportTemplatesPartialUpdateOK handles this case with default header values.

ExtrasExportTemplatesPartialUpdateOK extras export templates partial update o k
*/
type ExtrasExportTemplatesPartialUpdateOK struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/{id}/][%d] extrasExportTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesPartialUpdateOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesPartialUpdateDefault creates a ExtrasExportTemplatesPartialUpdateDefault with default headers values
func NewExtrasExportTemplatesPartialUpdateDefault(code int) *ExtrasExportTemplatesPartialUpdateDefault {
	return &ExtrasExportTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*ExtrasExportTemplatesPartialUpdateDefault handles this case with default header values.

ExtrasExportTemplatesPartialUpdateDefault extras export templates partial update default
*/
type ExtrasExportTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras export templates partial update default response
func (o *ExtrasExportTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasExportTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/{id}/][%d] extras_export-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
