package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var errorStruct errors

func loginPage(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	cookie, err := c.Request.Cookie("session")
	if err == nil && session[cookie.Value] != "" {
		c.Redirect(http.StatusSeeOther, "/home")
		return
	}
	if c.Request.Method == "POST" {
		username := c.Request.FormValue("userName")
		password := c.Request.FormValue("passWord")

		if _, ok := users[username]; !ok {
			errorStruct.UsernameError = "Invalid Username"
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
		if password != users[username].Password {
			errorStruct.PasswordError = "Invalid Password"
			c.Redirect(303, "/")
			return
		}
		if password == users[username].Password {

			uid := uuid.New().String()
			cookie := &http.Cookie{
				Name:  "session",
				Value: uid,
			}
			http.SetCookie(c.Writer, cookie)
			session[cookie.Value] = username
			c.Redirect(http.StatusSeeOther, "/home")
			return
		}
	}

	tmp.ExecuteTemplate(c.Writer, "login.html", errorStruct)
}
func home(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	cookie, err := c.Request.Cookie("session")
	if err != nil || session[cookie.Value] == "" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	username := session[cookie.Value]
	us := users[username]
	tmp.ExecuteTemplate(c.Writer, "home.html", us)
}
func logOut(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")
	c.SetCookie("session", "", -1, "", "", true, true)
	c.Redirect(http.StatusSeeOther, "/")
}
