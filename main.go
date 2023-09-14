package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

var tmp *template.Template

type UserSt struct {
	Username string
	Password string
}

var users = make(map[string]UserSt)
var session = make(map[string]string)

type errors struct {
	UsernameError string
	PasswordError string
}

func init() {
	tmp = template.Must(template.ParseGlob("templates/*.html"))

	users["akhu@gmail.com"] = UserSt{"Akhanya", "9552"}
	users["shandu@gmail.com"] = UserSt{"Shandrima", "9551"}
	users["akhil@gmail.com"] = UserSt{"Akhil", "9550"}

}
func main() {

	r := gin.Default()
	r.POST("/", loginPage)
	// r.POST("/", loginPage)
	r.GET("/home", home)
	r.GET("/logout", logOut)
	r.Run(":3000")
}
