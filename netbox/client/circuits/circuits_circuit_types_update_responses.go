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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mikeburke14/go-netbox/netbox/models"
)

// CircuitsCircuitTypesUpdateReader is a Reader for the CircuitsCircuitTypesUpdate structure.
type CircuitsCircuitTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTypesUpdateOK creates a CircuitsCircuitTypesUpdateOK with default headers values
func NewCircuitsCircuitTypesUpdateOK() *CircuitsCircuitTypesUpdateOK {
	return &CircuitsCircuitTypesUpdateOK{}
}

/*CircuitsCircuitTypesUpdateOK handles this case with default header values.

CircuitsCircuitTypesUpdateOK circuits circuit types update o k
*/
type CircuitsCircuitTypesUpdateOK struct {
	Payload *models.CircuitType
}

func (o *CircuitsCircuitTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTypesUpdateOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTypesUpdateDefault creates a CircuitsCircuitTypesUpdateDefault with default headers values
func NewCircuitsCircuitTypesUpdateDefault(code int) *CircuitsCircuitTypesUpdateDefault {
	return &CircuitsCircuitTypesUpdateDefault{
		_statusCode: code,
	}
}

/*CircuitsCircuitTypesUpdateDefault handles this case with default header values.

CircuitsCircuitTypesUpdateDefault circuits circuit types update default
*/
type CircuitsCircuitTypesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit types update default response
func (o *CircuitsCircuitTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTypesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/{id}/][%d] circuits_circuit-types_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTypesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
