package v1

import (
	"net/http"

	"github.com/danielkrainas/gobag/errcode"
)

const ErrorGroup = "kindermud.api.v1"

var (
	ErrorCodeRequestInvalid = errcode.Register(ErrorGroup, errcode.ErrorDescriptor{
		Value:          "REQUEST_INVALID",
		Message:        "request validation failed",
		Description:    "",
		HTTPStatusCode: http.StatusBadRequest,
	})
)
