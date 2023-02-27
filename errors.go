package wsa_lib_utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ErrorDetails struct {
	ErrorCode        string       `json:"errorCode,omitempty"`
	ErrorDesc        string       `json:"errorDesc,omitempty"`
	MessageAddlnInfo string       `json:"messageAddlnInfo,omitempty"`
	RequestId        string       `json:"requestId,omitempty"`
	RespCode         *gin.Context `json:"respCode,omitempty"`
}

// FormatErrors is the exposed function for generating errors.
func FormatErrors(code int, message *ErrorDetails) error {
	if message == nil {
		return fmt.Errorf("%v", code)
	}

	buf, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("json marshall error ")
	}

	return fmt.Errorf("%v", string(buf))
}
