package errors

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	codeBadRequest          = http.StatusBadRequest          // 请求错误的错误码
	codeNotFound            = http.StatusNotFound            // 找不到的错误码
	codeServerInternalError = http.StatusInternalServerError // 系统内部错误的错误码
	codeDBError             = 1100                           // 数据库错误的错误码
	codePageTokenInvalid    = 1200                           // 非法的分页令牌的错误码
)

// Error 是错误结构体
type Error struct {
	err  error
	code int32
}

// Error 返回错误信息
func (e *Error) Error() string {
	if e == nil || e.err == nil {
		return ""
	}

	return fmt.Sprintf("%d (%s)", e.code, e.err.Error())
}

// Is 判断 e 是否是 target 类型的错误
func (e *Error) Is(target error) bool {
	if e == nil {
		return e == target
	}

	err, ok := target.(*Error)
	if !ok {
		return e.err == target
	}

	return e.code == err.code
}

// Unwrap 返回内部的错误
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.err
}

// ErrorInfo 详细错误信息
func (e *Error) ErrorInfo() (int32, error) {
	return e.code, e.err
}

// WithCode 将 code 包装进 error
func WithCode(err error, code int32) error {
	if err == nil {
		return nil
	}

	return &Error{
		err:  err,
		code: code,
	}
}

// Is 判断 err 是否是 code 类型的错误
func Is(err error, code int32) bool {
	for {
		if err == nil {
			return false
		}

		e, ok := err.(*Error)
		if !ok {
			return false
		}

		if e.code == code {
			return true
		}

		err = errors.Unwrap(err)
	}
}

func GetRawError(err error) *Error {
	ee, ok := err.(*Error)
	if ok {
		return ee
	}
	return &Error{fmt.Errorf("UnKnow error : %v", err), 9999}
}

// BadRequest 返回请求错误
func BadRequest(err error) error {
	return WithCode(err, codeBadRequest)
}

// IsBadRequest 判断 err 是不是请求错误
func IsBadRequest(err error) bool {
	return Is(err, codeBadRequest)
}

// NotFound 返回找不到的错误
func NotFound(err error) error {
	return WithCode(err, codeNotFound)
}

// IsNotFound 判断 err 是不是找不到的错误
func IsNotFound(err error) bool {
	return Is(err, codeNotFound)
}

// ServerInternalError 返回系统内部错误
func ServerInternalError(err error) error {
	return WithCode(err, codeServerInternalError)
}

// IsServerInternalError 判断 err 是不是系统内部错误
func IsServerInternalError(err error) bool {
	return Is(err, codeServerInternalError)
}

// DBError 返回数据库错误
func DBError(err error) error {
	return WithCode(err, codeDBError)
}

// IsDBError 判断 err 是不是数据库错误
func IsDBError(err error) bool {
	return Is(err, codeDBError)
}

// PageTokenInvalid 返回数据库错误
func PageTokenInvalid(err error) error {
	return WithCode(err, codePageTokenInvalid)
}

// IsPageTokenInvalid 判断 err 是不是数据库错误
func IsPageTokenInvalid(err error) bool {
	return Is(err, codePageTokenInvalid)
}
