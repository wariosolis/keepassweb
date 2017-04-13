package controllers

import (
	"github.com/revel/revel"
	"keepassweb/app/models"
)

// Dashboard Controller anon user
type Dashboard struct {
	App
}

// Index page
func (c Dashboard) Index() revel.Result {
	var categories = []models.Category{}
	c.Txn.Where("user_id = ?", c.connected().ID).Find(&categories)
	for index, element := range categories {
		var p = []models.Passwd{}
		c.Txn.Model(&element).Related(&p)
		element.Passwd = p
		categories[index].Passwd = p
	}
	return c.Render(categories)
}
