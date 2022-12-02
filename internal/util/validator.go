package util

import (
  "errors"
  "strings"

  "github.com/go-playground/validator/v10"
)

var Validate = validator.New()
var (
  ErrorsValidate = errors.New("入参校验未通过")
)

func init() {
  Validate.SetTagName("v")
}

type ErrorResponse struct {
  FailedField string
  Tag         string
  Value       string
}

func ValidateStruct(t interface{}) (error, []*ErrorResponse) {
  var errorResponses []*ErrorResponse
  err := Validate.Struct(t)
  if err != nil {
    for _, err := range err.(validator.ValidationErrors) {
      var element ErrorResponse
      element.FailedField = strings.ToLower(err.StructNamespace())
      element.Tag = strings.ToLower(err.Tag())
      element.Value = strings.ToLower(err.Param())
      errorResponses = append(errorResponses, &element)
    }
  }
  if len(errorResponses) > 0 {
    return ErrorsValidate, errorResponses
  }
  return nil, errorResponses
}
