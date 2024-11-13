package controllers

import (
	"github.com/xhermitx/itty-bitty/internal/utils"
	"html/template"
	"net/http"
)

func ShowIndex(w http.ResponseWriter, _ *http.Request) {
	path := utils.GetTemplatePath("index.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
