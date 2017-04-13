package controllers

import (
	"keepassweb/app/models"
	"github.com/revel/revel"
)

// PublicApp without permissions
type PublicApp struct {
	GormController
}

// AddUser func
// This function add user variable to ViewArgs
func (c PublicApp) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.ViewArgs["user"] = user
	}
	return nil
}

// connected return connected user
func (c PublicApp) connected() *models.User {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.currentUser(username)
	}
	return nil
}

// Ger current user by email
func (c PublicApp) currentUser(email string) *models.User {
	var currentUser = &models.User{}
	c.Txn.Where("email = ?", email).First(&currentUser)
	if c.Txn.Error != nil {
		return nil
	}
	c.ViewArgs["currentController"] = c.Name
	c.ViewArgs["currentAction"] = c.MethodName
	return currentUser
}

// App user loggued
type App struct {
	PublicApp
}

// Check if user is connected
func (c App) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Validation.Required(user != nil).Key("Email").Message("Permissions required")
		//return c.Redirect(routes.Account.Login())
		return c.Redirect("/")
	}
	return nil
}
