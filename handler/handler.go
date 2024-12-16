package handler

import (
	"context"
	"net/http"

	"github.com/aminasadiam/Shop/view/home"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	home.Index().Render(context.Background(), w)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}
