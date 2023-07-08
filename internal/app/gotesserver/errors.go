package gotesserver

import "errors"

var (
    errWrongEmailOrPassword = errors.New("incorrect email or password")
    errFailedIssueToken     = errors.New("failed to issue a token")
    errSyntaxParam          = errors.New("wrong syntax param")
    //err     = errors.New("failed to issue a token")
)

// An error structure
type ErrorResponse struct {
    Code  int    `json:"code"`
    Error string `json:"error"`
}

// Constructor for a new ErrorResponse
func NewErrorResponse(code int, err error) *ErrorResponse {
    return &ErrorResponse{
        Code:  code,
        Error: err.Error(),
    }
}
