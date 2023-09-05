package local

import (
	apinto_module "github.com/eolinker/apinto-dashboard/module"
)

type tDriver struct {
}

func NewDriver() apinto_module.Driver {
	return &tDriver{}
}

func (d *tDriver) CreatePlugin(define interface{}) (apinto_module.Plugin, error) {
	return newTPlugin(define)
}
