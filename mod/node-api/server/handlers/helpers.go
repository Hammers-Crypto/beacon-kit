// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/berachain/beacon-kit/mod/node-api/server/types"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		hasValidationErrors := errors.As(err, &validationErrors)
		if hasValidationErrors {
			return nil
		}
		firstError := validationErrors[0]
		field := firstError.Field()
		value := firstError.Value()
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("Invalid %s: %s", field, value))
	}
	return nil
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var message any = http.StatusText(code)
	httpError := &echo.HTTPError{}
	if errors.As(err, &httpError) {
		code = httpError.Code
		message = httpError.Message
	}
	c.Logger().Error(err)
	response := &types.ErrorResponse{
		Code:    code,
		Message: message,
	}
	if jsonErr := c.JSON(code, response); jsonErr != nil {
		c.Logger().Error(jsonErr)
	}
}

func BindAndValidate[T any](c echo.Context) (*T, error) {
	t := new(T)
	if err := c.Bind(t); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(t); err != nil {
		return nil, err
	}
	return t, nil
}

func WrapData(nested any) types.DataResponse {
	return types.DataResponse{Data: nested}
}
