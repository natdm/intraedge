package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/natdm/intraedge/tutorials/day2/application/models"
)

type API struct {
	store models.Storage
	path  string
}

func New(path string, s models.Storage) (*API, error) {
	api := &API{store: s, path: path}
	return api, nil
}

type path struct {
	url, method string
}

func (a *API) pathf(url, method string) path {
	return path{a.path + url, method}
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := path{url: r.URL.Path, method: r.Method}
	switch p {
	case a.pathf("", http.MethodGet), a.pathf("/", http.MethodGet):
		a.fileHandler(w, r)
	case a.pathf("", http.MethodPost), a.pathf("/", http.MethodPost):
		a.saveHandler(w, r)
	case a.pathf("/list", http.MethodGet), a.pathf("/list/", http.MethodGet):
		a.listHandler(w, r)
	}
}

// curl -F "uploadfile=@/Users/nathan.hyland/butternut.sh" localhost:8080/file
func (a *API) saveHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		// error out
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer file.Close()

	if err := a.store.Save(handler.Filename, file); err != nil {
		// error out
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) listHandler(w http.ResponseWriter, r *http.Request) {
	// List files
	names, err := a.store.List()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(models.NamesResponse{Names: names}); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *API) fileHandler(w http.ResponseWriter, r *http.Request) {
	// File files
	fname := r.URL.Query().Get("name")
	if len(fname) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := a.store.File(fname, w); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func MiddlwareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\nMETHOD:\t%s\nPATH:\t%s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
