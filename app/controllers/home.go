package controllers

import (
	"github.com/revel/revel"
)

// Home Controller anon user
type Home struct {
	PublicApp
}

// Index page
func (c Home) Index() revel.Result {
	return c.Render()
}
