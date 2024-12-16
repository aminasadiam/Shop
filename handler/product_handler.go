package handler

import (
	"context"
	"net/http"

	"github.com/aminasadiam/Shop/view/product"
)

func ProductListHandler(w http.ResponseWriter, r *http.Request) {
	product.List().Render(context.Background(), w)
}

func ProductCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	product.Categories().Render(context.Background(), w)
}
