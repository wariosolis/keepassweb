package controllers

import (
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"keepassweb/app/models"
	"keepassweb/app/routes"
)

// Account Controller
type Account struct {
	PublicApp
}

// Login method
func (c Account) Login() revel.Result {
	return c.Render()
}

// LoginProccess with data from Login form
func (c Account) LoginProccess(Email string, Password string) revel.Result {
	c.Validation.Required(Email)
	c.Validation.Email(Email)
	c.Validation.Required(Password)
	var user models.User
	c.Txn.Where("email = ?", Email).First(&user)
	if c.Txn.Error != nil {
		panic(c.Txn.Error)
	}
	c.Validation.Required(user.ID != 0).Key("email").Message("Email or Password incorrect")
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(Password))
	c.Validation.Required(err == nil).Key("password").Message("Email or Password incorrect")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.Flash.Error("Login failed")
		if user.ID == 0 {
			return c.Redirect(routes.Account.Register())
		}
		return c.Redirect(routes.Account.Login())
	}
	c.Session["user"] = user.Email
	c.Flash.Success("Welcome, " + user.Email)
	return c.Redirect("/")
}

// Register new user
func (c Account) Register() revel.Result {
	return c.Render()
}

// RegisterPost is a post of information input on Register
func (c Account) RegisterPost(Email string, Name string, Password string, Password2 string) revel.Result {
	c.Validation.Required(Email)
	c.Validation.Email(Email)
	c.Validation.Required(Name)
	c.Validation.Required(Password)
	c.Validation.Required(Password2)
	c.Validation.Required(Password == Password2).Key("password").Message("Claves no validas")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		return c.Redirect(routes.Account.Register())
	}
	var user = models.User{Email: Email, Name: Name, Active: true}
	user.SetNewPassword(Password)
	c.Txn.Create(&user)
	c.Txn.Save(&user)
	return c.Redirect("/")
}

// Logout ot web app
func (c Account) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Home.Index())
}
