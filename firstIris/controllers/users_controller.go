//apparently the Iris example has a user_controller && a users_controller
package controllers

import (
	"../actions"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

//new one created by Iris w/ each req.
type UsersController struct {
	//auto-binded
	Ctx     iris.Context
	Action  actions.UserActions
	Session *sessions.Session
}

const userIDKey = "UserID"

func (c *UsersController) getCurrentUserID() int64 {
	userID, _ := c.Session.GetInt64Default(userIDKey, 0) //key, defaultValue
	return userID
}

//in application_controller.rb
func (c *UsersController) signedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *UsersController) signout() {
	c.Session.Destroy()
}

var registerStaticView = mvc.View{
	Name: "users/gateway.html",
	Data: iris.Map{"Title": "User Gateway"},
}

func (c *UsersController) GetRegister() mvc.Result {
	if c.signedIn() {
		c.signout()
	}

	return registerStaticView //default view?
}

func (c *UsersController) PostRegister() mvc.Result {
	//get inputs from form
	var (
		firstname = c.Ctx.FormValue("firstname")
		username  = c.Ctx.FormValue("username")
		password  = c.Ctx.FormValue("password")
	)

	u, err := c.Action.Create(password, datamodels.User{
		Username:  username,
		Firstname: firstname,
	})

	c.Session.Set(userIDKey, u.ID)

	return mvc.Response{
		Err:  err,
		Path: "/user/me",
	}
}
