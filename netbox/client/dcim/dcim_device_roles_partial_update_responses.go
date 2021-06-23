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

// DcimDeviceRolesPartialUpdateReader is a Reader for the DcimDeviceRolesPartialUpdate structure.
type DcimDeviceRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceRolesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceRolesPartialUpdateOK creates a DcimDeviceRolesPartialUpdateOK with default headers values
func NewDcimDeviceRolesPartialUpdateOK() *DcimDeviceRolesPartialUpdateOK {
	return &DcimDeviceRolesPartialUpdateOK{}
}

/*DcimDeviceRolesPartialUpdateOK handles this case with default header values.

DcimDeviceRolesPartialUpdateOK dcim device roles partial update o k
*/
type DcimDeviceRolesPartialUpdateOK struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-roles/{id}/][%d] dcimDeviceRolesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceRolesPartialUpdateOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceRolesPartialUpdateDefault creates a DcimDeviceRolesPartialUpdateDefault with default headers values
func NewDcimDeviceRolesPartialUpdateDefault(code int) *DcimDeviceRolesPartialUpdateDefault {
	return &DcimDeviceRolesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimDeviceRolesPartialUpdateDefault handles this case with default header values.

DcimDeviceRolesPartialUpdateDefault dcim device roles partial update default
*/
type DcimDeviceRolesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device roles partial update default response
func (o *DcimDeviceRolesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceRolesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-roles/{id}/][%d] dcim_device-roles_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
