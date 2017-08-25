package controllers

import (
	"net/http"
	"fmt"

	"github.com/yalay/OpenCC-Server/opencc"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
)

const (
	kTransTypeHK = "hk"
	kTransTypeTW = "tw"
)

const (
	s2hkConfigFile = "data/s2hk.json"
	s2twConfigFile = "data/s2twp.json"
)

var gRender *render.Render

func InitTemplate(tplPath string) {
	gRender = render.New(render.Options{
		Directory: tplPath,
	})
}

func S2tGetHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if gRender == nil {
		fmt.Fprintln(w, "OK. No translator template files.")
		return
	}

	gRender.HTML(w, http.StatusOK, "translator", nil)
}

func S2tPostHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if err := req.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	oriText := req.FormValue("input")
	if oriText == "" {
		http.Error(w, "empty input", http.StatusBadRequest)
		return
	}

	configFile := ""
	transType := req.FormValue("area")
	switch transType {
	case kTransTypeHK:
		configFile = s2hkConfigFile
	case kTransTypeTW:
		configFile = s2twConfigFile
	}

	if configFile == "" {
		http.Error(w, "translate type unsupport.", http.StatusBadRequest)
		return
	}

	converter := opencc.NewConverter(configFile)
	transText := converter.Convert(oriText)
	converter.Close()
	gRender.Text(w, http.StatusOK, transText)
	return
}
