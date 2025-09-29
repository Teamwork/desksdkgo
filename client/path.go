package client

import "strconv"

type DefaultPathHandler struct {
	base string
}

func NewDefaultPathHandler(base string) DefaultPathHandler {
	return DefaultPathHandler{base: base}
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
