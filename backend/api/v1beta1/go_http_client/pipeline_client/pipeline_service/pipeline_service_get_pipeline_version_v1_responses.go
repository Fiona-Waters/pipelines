// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	pipeline_model "github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/pipeline_model"
)

// PipelineServiceGetPipelineVersionV1Reader is a Reader for the PipelineServiceGetPipelineVersionV1 structure.
type PipelineServiceGetPipelineVersionV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PipelineServiceGetPipelineVersionV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPipelineServiceGetPipelineVersionV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPipelineServiceGetPipelineVersionV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPipelineServiceGetPipelineVersionV1OK creates a PipelineServiceGetPipelineVersionV1OK with default headers values
func NewPipelineServiceGetPipelineVersionV1OK() *PipelineServiceGetPipelineVersionV1OK {
	return &PipelineServiceGetPipelineVersionV1OK{}
}

/*PipelineServiceGetPipelineVersionV1OK handles this case with default header values.

A successful response.
*/
type PipelineServiceGetPipelineVersionV1OK struct {
	Payload *pipeline_model.APIPipelineVersion
}

func (o *PipelineServiceGetPipelineVersionV1OK) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/pipeline_versions/{version_id}][%d] pipelineServiceGetPipelineVersionV1OK  %+v", 200, o.Payload)
}

func (o *PipelineServiceGetPipelineVersionV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.APIPipelineVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineServiceGetPipelineVersionV1Default creates a PipelineServiceGetPipelineVersionV1Default with default headers values
func NewPipelineServiceGetPipelineVersionV1Default(code int) *PipelineServiceGetPipelineVersionV1Default {
	return &PipelineServiceGetPipelineVersionV1Default{
		_statusCode: code,
	}
}

/*PipelineServiceGetPipelineVersionV1Default handles this case with default header values.

An unexpected error response.
*/
type PipelineServiceGetPipelineVersionV1Default struct {
	_statusCode int

	Payload *pipeline_model.GatewayruntimeError
}

// Code gets the status code for the pipeline service get pipeline version v1 default response
func (o *PipelineServiceGetPipelineVersionV1Default) Code() int {
	return o._statusCode
}

func (o *PipelineServiceGetPipelineVersionV1Default) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/pipeline_versions/{version_id}][%d] PipelineService_GetPipelineVersionV1 default  %+v", o._statusCode, o.Payload)
}

func (o *PipelineServiceGetPipelineVersionV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}