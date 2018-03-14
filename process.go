package main

import (
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
)

func PostProcess(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	memItem, err := getSession(r)
	if err != nil {
		log.Infof(ctx, "Attempt to post tweet from logged out user")
		http.Error(w, "You must be logged in", http.StatusForbidden)
		return
	} else {
		var user User
		json.Unmarshal(memItem.Value, &user)
		log.Infof(ctx, user.UserName)

		message := r.FormValue("Message")

		post := Message{
			Msg:      message,
			Created:  time.Now(),
			Username: user.UserName,
		}
		userKey := datastore.NewKey(ctx, "User", user.UserName, 0, nil)
		key := datastore.NewIncompleteKey(ctx, "Post", userKey)
		_, err := datastore.Put(ctx, key, &post)
		if err != nil {
			log.Errorf(ctx, "error adding todo: %v", err)
			http.Error(w, err.Error(), 500)
			return
		}

	}
	time.Sleep(time.Millisecond * 500)
	http.Redirect(w, r, "/home", 302)

}

func ReplyProcess(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	memItem, err := getSession(r)
	if err != nil {
		log.Errorf(ctx, "error getting postss: %v", err)
		http.Redirect(w, r, "/login", 302)
	} else {
		var user User
		json.Unmarshal(memItem.Value, &user)
		log.Infof(ctx, user.UserName)

		reply := r.FormValue("Reply")

		post := Reply{
			Responce: reply,
			Created:  time.Now(),
			Username: user.UserName,
		}
		userKey := datastore.NewKey(ctx, "User", user.UserName, 0, nil)
		key := datastore.NewIncompleteKey(ctx, "Reply", userKey)
		_, err := datastore.Put(ctx, key, &post)
		if err != nil {
			log.Errorf(ctx, "error adding todo: %v", err)
			http.Error(w, err.Error(), 500)
			return
		}

	}
	time.Sleep(time.Millisecond * 500)
	http.Redirect(w, r, "/", 302)

}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	NewUser := User{
		Email:    r.FormValue("email"),
		UserName: r.FormValue("userName"),
		Password: r.FormValue("password"),
	}
	key := datastore.NewKey(ctx, "User", NewUser.UserName, 0, nil)
	key, err := datastore.Put(ctx, key, &NewUser)

	if err != nil {
		log.Errorf(ctx, "error: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}

	id := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "session",
		Value: id.String(),
	}
	http.SetCookie(w, cookie)
	json, err := json.Marshal(NewUser)
	if err != nil {
		log.Errorf(ctx, "error marshalling: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	sd := memcache.Item{
		Key:   id.String(),
		Value: json,
	}
	memcache.Set(ctx, &sd)

	item, _ := memcache.Get(ctx, cookie.Value)
	if item != nil {
		log.Infof(ctx, "%s", string(item.Value))
	}

	http.Redirect(w, r, "/", 302)

}

func loginProcess(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	key := datastore.NewKey(ctx, "User", r.FormValue("userName"), 0, nil)
	var user User
	err := datastore.Get(ctx, key, &user)
	if err != nil || r.FormValue("password") != user.Password {
		var sd SessionData
		sd.LoginFail = true
		tpl.ExecuteTemplate(w, "login.html", sd)
		return
	} else {
		user.UserName = r.FormValue("userName")
		createSession(w, r, user)
		tpl.ExecuteTemplate(w, "home.html", user)
	}
}
