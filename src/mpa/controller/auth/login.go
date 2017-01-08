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
	"net/http"
	"os"
)

var conf *oauth2.Config = &oauth2.Config{
	ClientID:     "80cfe9da0569d8ebf716",
	ClientSecret: "2aa53efd7b99c05055578f51c056dcbce1e53a81",
	Endpoint:     github.Endpoint,
}

type LoginController struct{}

func (controller *LoginController) ServeGet(ctx *route.Context) error {
	url := conf.AuthCodeURL("dummy-state", oauth2.AccessTypeOnline)
	ctx.ResponseWriter.Header().Set("Location", url)
	ctx.ResponseWriter.WriteHeader(http.StatusFound)
	return nil
}

type LoginCallbackController struct {
	UserDAO model.UserDAO
}

func (controller *LoginCallbackController) ServeGet(ctx *route.Context) error {
	params := ctx.Request.URL.Query()
	state := params["state"][0]
	code := params["code"][0]
	if state != "dummy-state" {
		ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return model.ErrInvalidToken
	}

	httpContext := context.Background()
	tok, err := conf.Exchange(httpContext, code)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(ctx.ResponseWriter, "Exchange failure: %s", err)
		return model.ErrInvalidToken
	}
	client := conf.Client(httpContext, tok)
	user, err := findOrCreateAuthenticatedUser(client, controller.UserDAO)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(ctx.ResponseWriter, "DB lookup failure: %s", err)
		return model.ErrInvalidToken
	}

	tokenString, err := model.CreateTokenString(user, []byte{1, 2, 3, 4})
	if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(ctx.ResponseWriter, "JWT encode failure: %s", err)
		return model.ErrInvalidToken
	}

	authCookie := &http.Cookie{
		Name:  "AUTH_TOKEN",
		Value: tokenString,
		Path:  "/",
	}
	http.SetCookie(ctx.ResponseWriter, authCookie)
	http.Redirect(ctx.ResponseWriter, ctx.Request, "/", http.StatusFound)
	return nil
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