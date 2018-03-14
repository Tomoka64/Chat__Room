package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", logoutProcess)
	http.HandleFunc("/api/createuser", CreateProcess)
	http.HandleFunc("/api/loginprocess", loginProcess)
	http.HandleFunc("/chatroom", Home)
	http.HandleFunc("/post", Post)
	http.HandleFunc("/post-process", PostProcess)
	http.HandleFunc("/reply-process", ReplyProcess)
	tpl = template.Must(template.ParseGlob("template/*"))
}
