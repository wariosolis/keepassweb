package controllers

import (
	"github.com/revel/revel"
)

type Category struct {
	PublicApp
}

func (c Category) Index() revel.Result {
	return c.Render()
}


