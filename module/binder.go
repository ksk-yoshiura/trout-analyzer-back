package module

import (
	"bytes"
	"io/ioutil"

	"github.com/labstack/echo"
)

type CustomBinder struct {
	binder echo.DefaultBinder
}

func NewBinder() *CustomBinder {
	return &CustomBinder{
		binder: echo.DefaultBinder{},
	}
}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	var b []byte
	if b, err = ioutil.ReadAll(c.Request().Body); err != nil {
		return
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(b))

	err = cb.binder.Bind(i, c)

	req := c.Request()
	req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	c.SetRequest(req)

	return
}
