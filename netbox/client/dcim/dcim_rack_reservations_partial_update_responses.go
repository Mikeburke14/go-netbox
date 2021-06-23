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

// DcimRackReservationsPartialUpdateReader is a Reader for the DcimRackReservationsPartialUpdate structure.
type DcimRackReservationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackReservationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackReservationsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackReservationsPartialUpdateOK creates a DcimRackReservationsPartialUpdateOK with default headers values
func NewDcimRackReservationsPartialUpdateOK() *DcimRackReservationsPartialUpdateOK {
	return &DcimRackReservationsPartialUpdateOK{}
}

/*DcimRackReservationsPartialUpdateOK handles this case with default header values.

DcimRackReservationsPartialUpdateOK dcim rack reservations partial update o k
*/
type DcimRackReservationsPartialUpdateOK struct {
	Payload *models.RackReservation
}

func (o *DcimRackReservationsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-reservations/{id}/][%d] dcimRackReservationsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackReservationsPartialUpdateOK) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackReservationsPartialUpdateDefault creates a DcimRackReservationsPartialUpdateDefault with default headers values
func NewDcimRackReservationsPartialUpdateDefault(code int) *DcimRackReservationsPartialUpdateDefault {
	return &DcimRackReservationsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimRackReservationsPartialUpdateDefault handles this case with default header values.

DcimRackReservationsPartialUpdateDefault dcim rack reservations partial update default
*/
type DcimRackReservationsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack reservations partial update default response
func (o *DcimRackReservationsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackReservationsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-reservations/{id}/][%d] dcim_rack-reservations_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackReservationsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackReservationsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
