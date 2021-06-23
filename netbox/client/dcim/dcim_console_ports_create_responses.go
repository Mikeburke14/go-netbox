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

// DcimConsolePortsCreateReader is a Reader for the DcimConsolePortsCreate structure.
type DcimConsolePortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimConsolePortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortsCreateCreated creates a DcimConsolePortsCreateCreated with default headers values
func NewDcimConsolePortsCreateCreated() *DcimConsolePortsCreateCreated {
	return &DcimConsolePortsCreateCreated{}
}

/*DcimConsolePortsCreateCreated handles this case with default header values.

DcimConsolePortsCreateCreated dcim console ports create created
*/
type DcimConsolePortsCreateCreated struct {
	Payload *models.ConsolePort
}

func (o *DcimConsolePortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/console-ports/][%d] dcimConsolePortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimConsolePortsCreateCreated) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortsCreateDefault creates a DcimConsolePortsCreateDefault with default headers values
func NewDcimConsolePortsCreateDefault(code int) *DcimConsolePortsCreateDefault {
	return &DcimConsolePortsCreateDefault{
		_statusCode: code,
	}
}

/*DcimConsolePortsCreateDefault handles this case with default header values.

DcimConsolePortsCreateDefault dcim console ports create default
*/
type DcimConsolePortsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console ports create default response
func (o *DcimConsolePortsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsolePortsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/console-ports/][%d] dcim_console-ports_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
