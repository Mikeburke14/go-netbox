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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mikeburke14/go-netbox/netbox/models"
)

// DcimConsolePortsPartialUpdateReader is a Reader for the DcimConsolePortsPartialUpdate structure.
type DcimConsolePortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortsPartialUpdateOK creates a DcimConsolePortsPartialUpdateOK with default headers values
func NewDcimConsolePortsPartialUpdateOK() *DcimConsolePortsPartialUpdateOK {
	return &DcimConsolePortsPartialUpdateOK{}
}

/*DcimConsolePortsPartialUpdateOK handles this case with default header values.

DcimConsolePortsPartialUpdateOK dcim console ports partial update o k
*/
type DcimConsolePortsPartialUpdateOK struct {
	Payload *models.ConsolePort
}

func (o *DcimConsolePortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-ports/{id}/][%d] dcimConsolePortsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortsPartialUpdateOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortsPartialUpdateDefault creates a DcimConsolePortsPartialUpdateDefault with default headers values
func NewDcimConsolePortsPartialUpdateDefault(code int) *DcimConsolePortsPartialUpdateDefault {
	return &DcimConsolePortsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimConsolePortsPartialUpdateDefault handles this case with default header values.

DcimConsolePortsPartialUpdateDefault dcim console ports partial update default
*/
type DcimConsolePortsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console ports partial update default response
func (o *DcimConsolePortsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsolePortsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-ports/{id}/][%d] dcim_console-ports_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
