package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	msgs := make([]*Message, 0, 10)
	q := datastore.NewQuery("Post").Order("-createdAt").Limit(cap(msgs))
	for it := q.Run(ctx); ; {
		var msg Message
		_, err := it.Next(&msg)
		if err == datastore.Done {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		msgs = append(msgs, &msg)
	}
	fmt.Println(ctx, msgs)

	if err := tpl.ExecuteTemplate(w, "public.html", msgs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	ServeTemplate(w, "login.html", r)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	ServeTemplate(w, "signup.html", r)
}

func Home(w http.ResponseWriter, r *http.Request) {
	ServeTemplate(w, "home.html", r)
}

func Post(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	memItem, err := getSession(r)
	if err != nil {
		log.Errorf(ctx, "error getting postss: %v", err)
		http.Redirect(w, r, "/login", 302)
	} else {
		var sd SessionData
		if err == nil {
			// logged in
			json.Unmarshal(memItem.Value, &sd)
			sd.LoggedIn = true
		}
		tpl.ExecuteTemplate(w, "post.html", &sd)
	}
}

func ServeTemplate(w http.ResponseWriter, templateName string, r *http.Request) {
	memItem, err := getSession(r)
	if err != nil {
		// not logged in
		tpl.ExecuteTemplate(w, templateName, SessionData{})
	} else {
		// logged in
		var sd SessionData
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
		tpl.ExecuteTemplate(w, templateName, sd)
	}
}
