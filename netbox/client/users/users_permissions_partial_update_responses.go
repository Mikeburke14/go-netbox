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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mikeburke14/go-netbox/netbox/models"
)

// UsersPermissionsPartialUpdateReader is a Reader for the UsersPermissionsPartialUpdate structure.
type UsersPermissionsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersPermissionsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersPermissionsPartialUpdateOK creates a UsersPermissionsPartialUpdateOK with default headers values
func NewUsersPermissionsPartialUpdateOK() *UsersPermissionsPartialUpdateOK {
	return &UsersPermissionsPartialUpdateOK{}
}

/*UsersPermissionsPartialUpdateOK handles this case with default header values.

UsersPermissionsPartialUpdateOK users permissions partial update o k
*/
type UsersPermissionsPartialUpdateOK struct {
	Payload *models.ObjectPermission
}

func (o *UsersPermissionsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/permissions/{id}/][%d] usersPermissionsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersPermissionsPartialUpdateOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersPermissionsPartialUpdateDefault creates a UsersPermissionsPartialUpdateDefault with default headers values
func NewUsersPermissionsPartialUpdateDefault(code int) *UsersPermissionsPartialUpdateDefault {
	return &UsersPermissionsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*UsersPermissionsPartialUpdateDefault handles this case with default header values.

UsersPermissionsPartialUpdateDefault users permissions partial update default
*/
type UsersPermissionsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users permissions partial update default response
func (o *UsersPermissionsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UsersPermissionsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /users/permissions/{id}/][%d] users_permissions_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersPermissionsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
