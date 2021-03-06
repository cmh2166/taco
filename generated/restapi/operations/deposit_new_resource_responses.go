// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DepositNewResourceOKCode is the HTTP code returned for type DepositNewResourceOK
const DepositNewResourceOKCode int = 200

/*DepositNewResourceOK OK

swagger:response depositNewResourceOK
*/
type DepositNewResourceOK struct {
}

// NewDepositNewResourceOK creates DepositNewResourceOK with default headers values
func NewDepositNewResourceOK() *DepositNewResourceOK {
	return &DepositNewResourceOK{}
}

// WriteResponse to the client
func (o *DepositNewResourceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DepositNewResourceMethodNotAllowedCode is the HTTP code returned for type DepositNewResourceMethodNotAllowed
const DepositNewResourceMethodNotAllowedCode int = 405

/*DepositNewResourceMethodNotAllowed Invalid input

swagger:response depositNewResourceMethodNotAllowed
*/
type DepositNewResourceMethodNotAllowed struct {
}

// NewDepositNewResourceMethodNotAllowed creates DepositNewResourceMethodNotAllowed with default headers values
func NewDepositNewResourceMethodNotAllowed() *DepositNewResourceMethodNotAllowed {
	return &DepositNewResourceMethodNotAllowed{}
}

// WriteResponse to the client
func (o *DepositNewResourceMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
