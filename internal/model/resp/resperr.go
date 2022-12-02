package resp

import "errors"

type IErrorCode interface {
  Error() string
  StatusCode() int
  Data() interface{}
}

type Error struct {
  code int
  err  error
  data interface{}
}

func (e *Error) Error() string {
  return e.err.Error()
}

func (e *Error) StatusCode() int {
  return e.code
}

func (e *Error) Data() interface{} {
  return e.data
}

func NewResperr(code int, e error, datas ...interface{}) *Error {
  var data interface{}
  if len(datas) > 0 {
    data = datas[0]
  }
  return &Error{
    code: code,
    err:  e,
    data: data,
  }
}
func NewResperrRawerr(err error) *Error {
  return &Error{
    code: ERROR,
    err:  err,
    data: nil,
  }
}

func NewResperrText(msg string) *Error {
  return &Error{
    code: ERROR,
    err:  errors.New(msg),
    data: nil,
  }
}
