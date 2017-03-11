package auth

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"mpa/model"
	"mpa/route"
	utilview "mpa/view/util"
	"net/http"
	"os"
)

var conf *oauth2.Config = &oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	Endpoint:     github.Endpoint,
}

type LoginController struct{}

// ServeGet implements route.GetController
func (controller *LoginController) ServeGet(ctx *route.Context) (route.View, error) {
	url := conf.AuthCodeURL("dummy-state", oauth2.AccessTypeOnline)
	return &utilview.RedirectView{
		Url: url,
	}, nil
}

type LoginCallbackController struct {
	UserDAO model.UserDAO
}

// ServeGet implements route.GetController
func (controller *LoginCallbackController) ServeGet(ctx *route.Context) (route.View, error) {
	params := ctx.Request.URL.Query()
	state := params["state"][0]
	code := params["code"][0]
	if state != "dummy-state" {
		ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return nil, model.ErrInvalidToken
	}

	httpContext := context.Background()
	tok, err := conf.Exchange(httpContext, code)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(ctx.ResponseWriter, "Exchange failure: %s", err)
		return nil, model.ErrInvalidToken
	}
	client := conf.Client(httpContext, tok)
	user, err := findOrCreateAuthenticatedUser(client, controller.UserDAO)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(ctx.ResponseWriter, "DB lookup failure: %s", err)
		return nil, model.ErrInvalidToken
	}

	tokenString, err := model.CreateTokenString(user, []byte{1, 2, 3, 4})
	if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(ctx.ResponseWriter, "JWT encode failure: %s", err)
		return nil, model.ErrInvalidToken
	}

	authCookie := &http.Cookie{
		Name:  "AUTH_TOKEN",
		Value: tokenString,
		Path:  "/",
	}
	http.SetCookie(ctx.ResponseWriter, authCookie)
	return &utilview.RedirectView{
		Url: "/",
	}, nil
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

	usr, err := userDAO.FindByLogin(githubUser.Login)
	if err == model.ErrNoEntry {
		usr = model.User{
			Login: githubUser.Login,
			Name:  githubUser.Name,
		}
		_, err = userDAO.Create(&usr)
		if err != nil {
			return model.User{}, nil
		}
		return usr, nil
	}
	return usr, err
}
