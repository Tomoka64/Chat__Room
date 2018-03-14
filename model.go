package main

import (
  "time"
  "google.golang.org/appengine/user"
)

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

type SessionData struct {
	User
	LoggedIn  bool
	LoginFail bool
	Message   Message
}

type Message struct {
  Isbn string `json:"isbn"`
  Username string`datastore:"username"`
  Created time.Time `datastore:"createdAt"`
  Msg string `datastore:"message"`
  Reply
}

type Reply struct {
  Isbn string
  Username user.User
  Responce string `datastore:"reply"`
  Created time.Time `datastore:"replycreatedAt"`
}
