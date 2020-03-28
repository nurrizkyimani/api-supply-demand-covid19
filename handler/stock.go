package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
)

// StockResource holds stock data store information.
type StockResource struct {
	*model.StockDataStore
}

func (store *StockResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(PaginationCtx)
	r.Get("/", GetAllStock(store))

	return r
}

// GetAllStock return stocks
func GetAllStock(repo interface {
	model.HasGetAllStock
}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paging, _ := r.Context().Value(PageCtxKey).(*Paging)

		stockDataPage, err := repo.GetAllStock(r.Context(), paging.Page, paging.Size)

		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, stockDataPage)
	}
}
