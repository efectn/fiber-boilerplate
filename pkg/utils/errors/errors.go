package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2/utils"
)

// Error represents an error that occurred while handling a request.
type Error struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

// Error makes it compatible with the `error` interface.
func (e *Error) Error() string {
	return fmt.Sprint(e.Message)
}

// NewErrors creates multiple new Error messages
func NewErrors(code int, messages ...interface{}) *Error {
	e := &Error{
		Code:    code,
		Message: utils.StatusMessage(code),
	}
	if len(messages) > 0 {
		e.Message = messages
	}
	return e
}
