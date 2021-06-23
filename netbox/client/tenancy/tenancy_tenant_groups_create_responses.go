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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mikeburke14/go-netbox/netbox/models"
)

// TenancyTenantGroupsCreateReader is a Reader for the TenancyTenantGroupsCreate structure.
type TenancyTenantGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyTenantGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantGroupsCreateCreated creates a TenancyTenantGroupsCreateCreated with default headers values
func NewTenancyTenantGroupsCreateCreated() *TenancyTenantGroupsCreateCreated {
	return &TenancyTenantGroupsCreateCreated{}
}

/*TenancyTenantGroupsCreateCreated handles this case with default header values.

TenancyTenantGroupsCreateCreated tenancy tenant groups create created
*/
type TenancyTenantGroupsCreateCreated struct {
	Payload *models.TenantGroup
}

func (o *TenancyTenantGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /tenancy/tenant-groups/][%d] tenancyTenantGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyTenantGroupsCreateCreated) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantGroupsCreateDefault creates a TenancyTenantGroupsCreateDefault with default headers values
func NewTenancyTenantGroupsCreateDefault(code int) *TenancyTenantGroupsCreateDefault {
	return &TenancyTenantGroupsCreateDefault{
		_statusCode: code,
	}
}

/*TenancyTenantGroupsCreateDefault handles this case with default header values.

TenancyTenantGroupsCreateDefault tenancy tenant groups create default
*/
type TenancyTenantGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy tenant groups create default response
func (o *TenancyTenantGroupsCreateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyTenantGroupsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /tenancy/tenant-groups/][%d] tenancy_tenant-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
