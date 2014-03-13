package controllers

import (
	"github.com/revel/revel"
	"problem-dashboard/app/models"
	"code.google.com/p/go.crypto/bcrypt"
	"problem-dashboard/app/routes"
)

type App struct {
	GorpController
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) SaveUser(user models.User, verifyPassword string) revel.Result {
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).
		Message("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Register())
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)

	err := c.Txn.Insert(&user)
	if err != nil {
		panic(err)
	}

	c.Session["user"] = user.UserName
	c.Flash.Success("Welcome, " + user.Name)
	return c.Redirect(routes.Problems.Index())
}
