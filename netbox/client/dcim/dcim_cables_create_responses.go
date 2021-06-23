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

// DcimCablesCreateReader is a Reader for the DcimCablesCreate structure.
type DcimCablesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimCablesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCablesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCablesCreateCreated creates a DcimCablesCreateCreated with default headers values
func NewDcimCablesCreateCreated() *DcimCablesCreateCreated {
	return &DcimCablesCreateCreated{}
}

/*DcimCablesCreateCreated handles this case with default header values.

DcimCablesCreateCreated dcim cables create created
*/
type DcimCablesCreateCreated struct {
	Payload *models.Cable
}

func (o *DcimCablesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/cables/][%d] dcimCablesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimCablesCreateCreated) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCablesCreateDefault creates a DcimCablesCreateDefault with default headers values
func NewDcimCablesCreateDefault(code int) *DcimCablesCreateDefault {
	return &DcimCablesCreateDefault{
		_statusCode: code,
	}
}

/*DcimCablesCreateDefault handles this case with default header values.

DcimCablesCreateDefault dcim cables create default
*/
type DcimCablesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cables create default response
func (o *DcimCablesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimCablesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/cables/][%d] dcim_cables_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
