package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io"
	"log"
	"mpa/auth"
	"mpa/handler"
	"mpa/model"
	"mpa/route"
	"net/http"
	"os"
)

func registerAPI(resource string, showHandler, searchHandler http.Handler) {
	http.Handle("/api/"+resource+"/", showHandler)
	http.Handle("/api/"+resource, searchHandler)
}

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
		log.Fatalf("Usage: %s MYSQL_SERVER_ADDRESS", os.Args[0])
	}
	addr := os.Args[1]

	db := sqlx.MustConnect("mysql", "mpa@tcp("+addr+":3306)/mpa")
	pluginDAO := model.NewPluginMySQLDAO(db)
	userDAO := model.NewUserMySQLDAO(db)
	sessionDAO := model.NewSessionMySQLDAO(db, userDAO)

	authFilterChain := route.CreateFilterChain(&auth.Filter{[]byte{1, 2, 3, 4}})
	registerAPI("plugin", handler.NewPluginHandler(pluginDAO), handler.NewPluginSearchHandler(pluginDAO))
	http.Handle("/api/user", authFilterChain.Wrap(mainPageHandler))
	http.HandleFunc("/api/auth/login", auth.LoginHandler)
	http.HandleFunc("/api/auth/callback", auth.LoginCallbackHandler(userDAO, sessionDAO))
	http.HandleFunc("/static/", staticFileHandler)
	http.HandleFunc("/", mainPageHandler)
	http.ListenAndServe(":8080", nil)
}
