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

// UsersPermissionsCreateReader is a Reader for the UsersPermissionsCreate structure.
type UsersPermissionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersPermissionsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersPermissionsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersPermissionsCreateCreated creates a UsersPermissionsCreateCreated with default headers values
func NewUsersPermissionsCreateCreated() *UsersPermissionsCreateCreated {
	return &UsersPermissionsCreateCreated{}
}

/*UsersPermissionsCreateCreated handles this case with default header values.

UsersPermissionsCreateCreated users permissions create created
*/
type UsersPermissionsCreateCreated struct {
	Payload *models.ObjectPermission
}

func (o *UsersPermissionsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /users/permissions/][%d] usersPermissionsCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersPermissionsCreateCreated) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersPermissionsCreateDefault creates a UsersPermissionsCreateDefault with default headers values
func NewUsersPermissionsCreateDefault(code int) *UsersPermissionsCreateDefault {
	return &UsersPermissionsCreateDefault{
		_statusCode: code,
	}
}

/*UsersPermissionsCreateDefault handles this case with default header values.

UsersPermissionsCreateDefault users permissions create default
*/
type UsersPermissionsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users permissions create default response
func (o *UsersPermissionsCreateDefault) Code() int {
	return o._statusCode
}

func (o *UsersPermissionsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /users/permissions/][%d] users_permissions_create default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersPermissionsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
