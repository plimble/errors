package errors

import (
	"errors"
	"fmt"

	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var mapGrpcCode = map[int]codes.Code{
	400: 3,
	401: 16,
	403: 7,
	404: 5,
	441: 4,
	500: 13,
	501: 12,
	503: 14,
	520: 2,
}

type HTTPError struct {
	status  int
	Message string `json:"error"`
	cause   error
}

func (e *HTTPError) Error() string { return e.Message }
func (e *HTTPError) Cause() error  { return e.cause }
func (e *HTTPError) WithCause(err error) *HTTPError {
	e.cause = err
	return e
}

func (e *HTTPError) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("error", e.Message)
	if e.cause != nil {
		enc.AddString("error_cause", e.cause.Error())
	}

	return nil
}

func Wrap(err error, httpError *HTTPError) error {
	if err != nil {
		return nil
	}

	return httpError.WithCause(err)
}

func New(msg string) error {
	return errors.New(msg)
}

func Newf(format string, v ...interface{}) error {
	return errors.New(fmt.Sprintf(format, v...))
}

func Error(status int, msg string) *HTTPError {
	return &HTTPError{status, msg, nil}
}

func Errorf(status int, format string, v ...interface{}) *HTTPError {
	return &HTTPError{status, fmt.Sprintf(format, v...), nil}
}

func BadRequest(msg string) *HTTPError {
	return &HTTPError{400, msg, nil}
}

func BadRequestf(format string, v ...interface{}) *HTTPError {
	return &HTTPError{400, fmt.Sprintf(format, v...), nil}
}

func Unauthorized(msg string) *HTTPError {
	return &HTTPError{401, msg, nil}
}

func Unauthorizedf(format string, v ...interface{}) *HTTPError {
	return &HTTPError{401, fmt.Sprintf(format, v...), nil}
}

func Forbidden(msg string) *HTTPError {
	return &HTTPError{403, msg, nil}
}

func Forbiddenf(format string, v ...interface{}) *HTTPError {
	return &HTTPError{403, fmt.Sprintf(format, v...), nil}
}

func NotFound(msg string) *HTTPError {
	return &HTTPError{404, msg, nil}
}

func NotFoundf(format string, v ...interface{}) *HTTPError {
	return &HTTPError{404, fmt.Sprintf(format, v...), nil}
}

func InternalError(msg string) *HTTPError {
	return &HTTPError{500, msg, nil}
}

func InternalErrorf(format string, v ...interface{}) *HTTPError {
	return &HTTPError{500, fmt.Sprintf(format, v...), nil}
}

func Timeout(msg string) *HTTPError {
	return &HTTPError{441, msg, nil}
}

func Timeoutf(format string, v ...interface{}) *HTTPError {
	return &HTTPError{441, fmt.Sprintf(format, v...), nil}
}

func NotImplement(msg string) *HTTPError {
	return &HTTPError{501, msg, nil}
}

func NotImplementf(format string, v ...interface{}) *HTTPError {
	return &HTTPError{501, fmt.Sprintf(format, v...), nil}
}

func Unavailable(msg string) *HTTPError {
	return &HTTPError{503, msg, nil}
}

func Unavailablef(format string, v ...interface{}) *HTTPError {
	return &HTTPError{503, fmt.Sprintf(format, v...), nil}
}

func UnknownError(msg string) *HTTPError {
	return &HTTPError{520, msg, nil}
}

func UnknownErrorf(format string, v ...interface{}) *HTTPError {
	return &HTTPError{520, fmt.Sprintf(format, v...), nil}
}

func Cause(err error) error {
	herr, ok := err.(*HTTPError)
	if !ok {
		return err
	}

	return herr.Cause()
}

func FromError(err error) (*HTTPError, bool) {
	herr, ok := err.(*HTTPError)
	return herr, ok
}

func ToGRPC(err error) error {
	if err == nil {
		return nil
	}

	herr, ok := err.(*HTTPError)
	if !ok {
		return status.Error(2, err.Error())
	}

	c := mapGrpcCode[herr.status]
	return status.Error(c, herr.Message)
}

func ErrStatus(err error) int {
	herr, ok := err.(*HTTPError)
	if !ok {
		return 0
	}

	return herr.status
}

func IsNotFound(err error) bool {
	return err != nil && ErrStatus(err) == 404
}

func IsInternalError(err error) bool {
	return err != nil && ErrStatus(err) == 500
}

func IsBadRequest(err error) bool {
	return err != nil && ErrStatus(err) == 400
}

func IsUnauthorized(err error) bool {
	return err != nil && ErrStatus(err) == 401
}

func IsForbidden(err error) bool {
	return err != nil && ErrStatus(err) == 403
}
