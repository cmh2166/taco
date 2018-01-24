// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DepositNewFileOKCode is the HTTP code returned for type DepositNewFileOK
const DepositNewFileOKCode int = 200

/*DepositNewFileOK OK

swagger:response depositNewFileOK
*/
type DepositNewFileOK struct {
}

// NewDepositNewFileOK creates DepositNewFileOK with default headers values
func NewDepositNewFileOK() *DepositNewFileOK {
	return &DepositNewFileOK{}
}

// WriteResponse to the client
func (o *DepositNewFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DepositNewFileMethodNotAllowedCode is the HTTP code returned for type DepositNewFileMethodNotAllowed
const DepositNewFileMethodNotAllowedCode int = 405

/*DepositNewFileMethodNotAllowed Invalid input

swagger:response depositNewFileMethodNotAllowed
*/
type DepositNewFileMethodNotAllowed struct {
}

// NewDepositNewFileMethodNotAllowed creates DepositNewFileMethodNotAllowed with default headers values
func NewDepositNewFileMethodNotAllowed() *DepositNewFileMethodNotAllowed {
	return &DepositNewFileMethodNotAllowed{}
}

// WriteResponse to the client
func (o *DepositNewFileMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}