package client

import (
	"net/http"
	"strconv"
)

type DefaultPathHandler struct {
	base         string
	updateMethod string
}

func NewDefaultPathHandler(base string) DefaultPathHandler {
	return NewDefaultPathHandlerWithUpdateMethod(base, http.MethodPut)
}

func NewDefaultPathHandlerWithUpdateMethod(base, updateMethod string) DefaultPathHandler {
	if updateMethod == "" {
		updateMethod = http.MethodPut
	}

	return DefaultPathHandler{base: base, updateMethod: updateMethod}
}

func (d DefaultPathHandler) Get(id int) string {
	return d.base + "/" + strconv.Itoa(id)
}

func (d DefaultPathHandler) List() string {
	return d.base
}

func (d DefaultPathHandler) Create() string {
	return d.base
}

func (d DefaultPathHandler) Update(id int) string {
	return d.base + "/" + strconv.Itoa(id)
}

func (d DefaultPathHandler) UpdateMethod() string {
	if d.updateMethod == "" {
		return http.MethodPut
	}

	return d.updateMethod
}
