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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mikeburke14/go-netbox/netbox/models"
)

// ExtrasImageAttachmentsPartialUpdateReader is a Reader for the ExtrasImageAttachmentsPartialUpdate structure.
type ExtrasImageAttachmentsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasImageAttachmentsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsPartialUpdateOK creates a ExtrasImageAttachmentsPartialUpdateOK with default headers values
func NewExtrasImageAttachmentsPartialUpdateOK() *ExtrasImageAttachmentsPartialUpdateOK {
	return &ExtrasImageAttachmentsPartialUpdateOK{}
}

/*ExtrasImageAttachmentsPartialUpdateOK handles this case with default header values.

ExtrasImageAttachmentsPartialUpdateOK extras image attachments partial update o k
*/
type ExtrasImageAttachmentsPartialUpdateOK struct {
	Payload *models.ImageAttachment
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extrasImageAttachmentsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsPartialUpdateDefault creates a ExtrasImageAttachmentsPartialUpdateDefault with default headers values
func NewExtrasImageAttachmentsPartialUpdateDefault(code int) *ExtrasImageAttachmentsPartialUpdateDefault {
	return &ExtrasImageAttachmentsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*ExtrasImageAttachmentsPartialUpdateDefault handles this case with default header values.

ExtrasImageAttachmentsPartialUpdateDefault extras image attachments partial update default
*/
type ExtrasImageAttachmentsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras image attachments partial update default response
func (o *ExtrasImageAttachmentsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extras_image-attachments_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
