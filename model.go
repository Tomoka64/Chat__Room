package main

import (
  "time"
)

type User struct {
	Email    string
	UserName string `datastore:"username"`
	Password string `json:"-"`
}

type SessionData struct {
	User
	LoggedIn  bool
	LoginFail bool
	Message   Message
}

type Message struct {
  Username string`datastore:"username"`
  Created time.Time `datastore:"createdAt"`
  Msg string `datastore:"message"`
  Reply
}

type Reply struct {
  Username string`datastore:"username"`
  Responce string `datastore:"reply"`
  Created time.Time `datastore:"createdAt"`
}
