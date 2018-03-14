package main

import (
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
)

func getSession(req *http.Request) (*memcache.Item, error) {
	cookie, err := req.Cookie("session")
	if err != nil {
		return &memcache.Item{}, err
	}
	ctx := appengine.NewContext(req)
	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
		return &memcache.Item{}, err
	}
	return item, nil
}

func createSession(w http.ResponseWriter, r *http.Request, user User) {
	ctx := appengine.NewContext(r)

	id := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "session",
		Value: id.String(),
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	json, err := json.Marshal(user)
	if err != nil {
		log.Errorf(ctx, "error marshalling: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	sd := memcache.Item{
		Key:        id.String(),
		Value:      json,
		Expiration: time.Duration(20 * time.Minute),
	}
	memcache.Set(ctx, &sd)
}
