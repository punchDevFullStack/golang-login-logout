package controllers

import (
	"fmt"

	"github.com/thanamat/golang-login-logout/models"
	"github.com/thanamat/golang-login-logout/utils"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	// start session
	sess := utils.GetSession()
	session, err := sess.SessionStart(l.Ctx.ResponseWriter, l.Ctx.Request)
	if err != nil {
		fmt.Println("session", err)
	}
	user := session.Get("user")

	fmt.Println("session user", user)

	if user != nil {
		// m := user.(map[models.Member]interface{})

		// fmt.Println("session user2", m)
		// l.Data["Title"] = email
		// user.(&models.Member)
		l.TplName = "index.html"
		return
	}

	l.TplName = "login.html"
	l.Render()
}

func (l *LoginController) Login() {

	email := l.GetString("email")
	password := l.GetString("password")

	fmt.Println(email, password)

	member, err := models.CheckMemberLogin(email, password)
	fmt.Println(member, err)
	if member == nil || err != nil {
		l.TplName = "login.html"
		l.Render()
		return
	}

	// start session
	sess := utils.GetSession()
	session2, err2 := sess.SessionStart(l.Ctx.ResponseWriter, l.Ctx.Request)
	if err2 != nil {
		fmt.Println("session", err2)
	}
	defer session2.SessionRelease(l.Ctx.ResponseWriter)
	session2.Set("user", member)
	l.Data["Title"] = email
	l.TplName = "index.html"

}

func (l *LoginController) Logout() {
	sess := utils.GetSession()

	sess.SessionDestroy(l.Ctx.ResponseWriter, l.Ctx.Request)

	l.Redirect("/login", 302)
}
