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

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mikeburke14/go-netbox/netbox/models"
)

// SecretsSecretRolesCreateReader is a Reader for the SecretsSecretRolesCreate structure.
type SecretsSecretRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecretsSecretRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecretsSecretRolesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecretsSecretRolesCreateCreated creates a SecretsSecretRolesCreateCreated with default headers values
func NewSecretsSecretRolesCreateCreated() *SecretsSecretRolesCreateCreated {
	return &SecretsSecretRolesCreateCreated{}
}

/*SecretsSecretRolesCreateCreated handles this case with default header values.

SecretsSecretRolesCreateCreated secrets secret roles create created
*/
type SecretsSecretRolesCreateCreated struct {
	Payload *models.SecretRole
}

func (o *SecretsSecretRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /secrets/secret-roles/][%d] secretsSecretRolesCreateCreated  %+v", 201, o.Payload)
}

func (o *SecretsSecretRolesCreateCreated) GetPayload() *models.SecretRole {
	return o.Payload
}

func (o *SecretsSecretRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsSecretRolesCreateDefault creates a SecretsSecretRolesCreateDefault with default headers values
func NewSecretsSecretRolesCreateDefault(code int) *SecretsSecretRolesCreateDefault {
	return &SecretsSecretRolesCreateDefault{
		_statusCode: code,
	}
}

/*SecretsSecretRolesCreateDefault handles this case with default header values.

SecretsSecretRolesCreateDefault secrets secret roles create default
*/
type SecretsSecretRolesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the secrets secret roles create default response
func (o *SecretsSecretRolesCreateDefault) Code() int {
	return o._statusCode
}

func (o *SecretsSecretRolesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /secrets/secret-roles/][%d] secrets_secret-roles_create default  %+v", o._statusCode, o.Payload)
}

func (o *SecretsSecretRolesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *SecretsSecretRolesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
