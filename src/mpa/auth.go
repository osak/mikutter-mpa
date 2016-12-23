package main

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
)

var conf *oauth2.Config = &oauth2.Config{
	ClientID:     "80cfe9da0569d8ebf716",
	ClientSecret: "2aa53efd7b99c05055578f51c056dcbce1e53a81",
	Endpoint:     github.Endpoint,
}

func authenticationStartHandler(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL("dummy-state", oauth2.AccessTypeOnline)
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}

func authenticationCallbackHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	state := params["state"][0]
	code := params["code"][0]
	if state != "dummy-state" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := conf.Exchange(context.Background(), code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusFound)
}
