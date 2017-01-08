package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"io"
	"log"
	"mpa/controller/auth"
	"mpa/controller/plugin"
	"mpa/controller/user"
	"mpa/filter"
	"mpa/model"
	"mpa/route"
	"net/http"
	"os"
)

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	f, err := os.Open("/app/web/" + path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not find %s\n", path)
		return
	}
	defer f.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.Copy(w, f)
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("/app/web/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not find index.html\n")
		return
	}
	defer f.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.Copy(w, f)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s MONGO_SERVER_ADDRESS", os.Args[0])
	}
	addr := os.Args[1]

	router := route.NewRouter()
	session, err := mgo.Dial(addr)
	if err != nil {
		panic("Cannot connect mongo: " + err.Error())
	}
	db := session.DB("mpa")
	pluginDAO := &model.MongoPluginDAO{db.C("plugins")}
	userDAO := &model.MongoUserDAO{db.C("users")}
	tokenDecoder := &model.TokenDecoder{userDAO}

	pluginController := &plugin.PluginController{pluginDAO}
	pluginEntryController := &plugin.PluginEntryController{pluginDAO}
	loginController := &auth.LoginController{}
	loginCallbackController := &auth.LoginCallbackController{userDAO}
	currentUserController := &user.CurrentUserController{}

	authFilterChain := route.CreateFilterChain(&filter.LoginFilter{tokenDecoder, []byte{1, 2, 3, 4}})
	router.RegisterGet("/api/plugin/", pluginEntryController)
	router.RegisterGet("/api/plugin", pluginController)
	router.RegisterPost("/api/plugin", authFilterChain.WrapPost(pluginController))
	router.RegisterGet("/api/me", authFilterChain.WrapGet(currentUserController))
	router.RegisterGet("/api/auth/login", loginController)
	router.RegisterGet("/api/auth/callback", loginCallbackController)
	http.HandleFunc("/static/", staticFileHandler)
	http.HandleFunc("/", mainPageHandler)
	http.ListenAndServe(":8080", nil)
}
