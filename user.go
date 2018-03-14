package main

//
// func checkUserName(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	ctx := appengine.NewContext(r)
// 	bs, err := ioutil.ReadAll(r.Body)
// 	sbs := string(bs)
// 	log.Infof(ctx, "REQUEST BODY:%v", sbs)
// 	var user User
// 	key := datastore.NewKey(ctx, "Users", sbs, 0, nil)
// 	err = datastore.Get(ctx, key, &user)
// 	log.Infof(ctx, "ERR: %v", err)
// 	if err != nil {
// 		fmt.Fprint(w, "false")
// 	} else {
// 		fmt.Fprint(w, "true")
// 	}
// }
//
// func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	ctx := appengine.NewContext(r)
// 	NewUser := User{
// 		Email:    r.FormValue("email"),
// 		UserName: r.FormValue("userName"),
// 		Password: r.FormValue("password"),
// 	}
// 	key := datastore.NewKey(ctx, "Users", NewUser.UserName, 0, nil)
// 	key, err := datastore.Put(ctx, key, &NewUser)
//
// 	if err != nil {
// 		log.Errorf(ctx, "error adding todo: %v", err)
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
//
// 	id := uuid.NewV4()
// 	cookie := &http.Cookie{
// 		Name:  "session",
// 		Value: id.String(),
// 		Path:  "/",
// 	}
// 	http.SetCookie(w, cookie)
// 	json, err := json.Marshal(NewUser)
// 	if err != nil {
// 		log.Errorf(ctx, "error marshalling during user creation: %v", err)
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	sd := memcache.Item{
// 		Key:   id.String(),
// 		Value: json,
// 	}
// 	memcache.Set(ctx, &sd)
// 	// TEST memcache
// 	item, _ := memcache.Get(ctx, cookie.Value)
// 	if item != nil {
// 		log.Infof(ctx, "%s", string(item.Value))
// 	}
//
// 	// redirect
// 	http.Redirect(w, r, "/", 302)
//
// }

//
// func loginProcess(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
// 	ctx := appengine.NewContext(req)
// 	key := datastore.NewKey(ctx, "Users", req.FormValue("userName"), 0, nil)
// 	var user User
// 	err := datastore.Get(ctx, key, &user)
// 	if err != nil || req.FormValue("password") != user.Password {
// 		var sd SessionData
// 		sd.LoginFail = true
// 		tpl.ExecuteTemplate(w, "login.html", sd)
// 		return
// 	} else {
// 		user.UserName = req.FormValue("userName")
// 		createSession(w, req, user)
// 		http.Redirect(w, req, "/", 302)
// 	}
// }
//
// func createSession(w http.ResponseWriter, r *http.Request, user User) {
// 	ctx := appengine.NewContext(r)
// 	//set cookie
// 	id := uuid.NewV4()
// 	cookie := &http.Cookie{
// 		Name:  "session",
// 		Value: id.String(),
// 		Path:  "/",
// 	}
// 	http.SetCookie(w, cookie)
//
// 	//set memcache session data(sd)
// 	json, err := json.Marshal(user)
// 	if err != nil {
// 		log.Errorf(ctx, "error marshalling during user creation: %v", err)
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	sd := memcache.Item{
// 		Key:        id.String(),
// 		Value:      json,
// 		Expiration: time.Duration(20 * time.Minute),
// 	}
// 	memcache.Set(ctx, &sd)
// }
//
// func logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	ctx := appengine.NewContext(r)
// 	cookie, err := r.Cookie("session")
// 	if err != nil {
// 		http.Redirect(w, r, "/", 302)
// 	}
//
// 	//clear memcache
// 	sd := memcache.Item{
// 		Key:        cookie.Value,
// 		Value:      []byte(""),
// 		Expiration: time.Duration(1 * time.Microsecond),
// 	}
// 	memcache.Set(ctx, &sd)
//
// 	//clear the Cookie
// 	cookie.MaxAge = -1
// 	http.SetCookie(w, cookie)
// 	http.Redirect(w, r, "/", 302)
// }
