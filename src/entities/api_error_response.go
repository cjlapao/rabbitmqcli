package entities

import (
	"fmt"

	"github.com/cjlapao/common-go/log"
)

type ApiErrorResponse struct {
	Code    ApiResponseErrorCode `json:"error_code"`
	Message string               `json:"error_message,omitempty"`
	Url     string               `json:"url,omitempty"`
}

func NewApiErrorResponse(code ApiResponseErrorCode, message string, url string) ApiErrorResponse {
	return ApiErrorResponse{
		Code:    code,
		Message: message,
		Url:     url,
	}
}

func (r ApiErrorResponse) Log() {
	logger := log.Get()
	msg := fmt.Sprintf("There was an error %s processing request", r.Code)
	if r.Url != "" {
		msg += fmt.Sprintf(" from %s", r.Url)
	}
	if r.Message != "" {
		msg += fmt.Sprintf(", %s", r.Message)
	}

	logger.Error(msg)
}

type ApiResponseErrorCode string

const (
	InvalidConnectionString ApiResponseErrorCode = "invalid_connection_string"
	InvalidClient           ApiResponseErrorCode = "invalid_client"
	BadRequest              ApiResponseErrorCode = "bad_request"
)
