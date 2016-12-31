package auth

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"mpa/model"
	"net/http"
	"os"
)

var conf *oauth2.Config = &oauth2.Config{
	ClientID:     "80cfe9da0569d8ebf716",
	ClientSecret: "2aa53efd7b99c05055578f51c056dcbce1e53a81",
	Endpoint:     github.Endpoint,
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL("dummy-state", oauth2.AccessTypeOnline)
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}

func LoginCallbackHandler(userDAO model.UserDAO, sessionDAO model.SessionDAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		state := params["state"][0]
		code := params["code"][0]
		if state != "dummy-state" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := context.Background()
		tok, err := conf.Exchange(ctx, code)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Exchange failure: %s", err)
			return
		}
		client := conf.Client(ctx, tok)
		user, err := findOrCreateAuthenticatedUser(client, userDAO)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB lookup failure: %s", err)
			return
		}

		tokenString, err := createTokenString(&user, []byte{1, 2, 3, 4})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "JWT encode failure: %s", err)
			return
		}
		w.Header().Set("X-JWT", tokenString)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

type githubUser struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

func findOrCreateAuthenticatedUser(client *http.Client, userDAO model.UserDAO) (model.User, error) {
	response, err := client.Get("https://api.github.com/user")
	if err != nil {
		return model.User{}, err
	}
	dec := json.NewDecoder(io.TeeReader(response.Body, os.Stdout))
	githubUser := githubUser{}
	err = dec.Decode(&githubUser)
	if err != nil {
		return model.User{}, err
	}

	user, err := userDAO.FindByLogin(githubUser.Login)
	if err == model.ErrNoEntry {
		user, err = userDAO.Create(model.User{
			Login: githubUser.Login,
			Name:  githubUser.Name,
		})
	}
	return user, err
}
