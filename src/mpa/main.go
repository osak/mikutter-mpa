package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"handler"
	"io"
	"log"
	"model"
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
		log.Fatalf("Usage: %s MYSQL_SERVER_ADDRESS", os.Args[0])
	}
	addr := os.Args[1]

	db := sqlx.MustConnect("mysql", "mpa@tcp("+addr+":3306)/mpa")
	dao := model.NewPluginMySQLDAO(db)

	http.Handle("/plugin", handler.NewPluginHandler(dao))
	http.HandleFunc("/static/", staticFileHandler)
	http.HandleFunc("/", mainPageHandler)
	http.ListenAndServe(":8080", nil)
}
