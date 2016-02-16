package main

import (
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/image/import", ImageImport)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if ctx == nil {
		http.Error(w, "Could not create app engine context", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Go-Scheduler")
}

func ImageImport(w http.ResponseWriter, r *http.Request) {
	resp, err := HttpGetter("http://admin.curtmfg.com/images/import", w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Image Import Response Code:"+resp.Status)
}

func InstallSheetImport(w http.ResponseWriter, r *http.Request) {
	resp, err := HttpGetter("http://admin.curtmfg.com/ins/import", w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "INS Import Response Code:"+resp.Status)
}

func HttpGetter(url string, w http.ResponseWriter, r *http.Request) (resp *http.Response, err error) {
	ctx := appengine.NewContext(r)
	if ctx == nil {
		err = errors.New("Could not create app engine context")
		return resp, err
	}
	client := urlfetch.Client(ctx)
	resp, err = client.Get(url)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
