package handler

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
)

// API provides application resources and handlers.
type API struct {
	authResource    *AuthResource
	stockResource   *StockResource
	requestResource *RequestResource
	itemResource *ItemResource
}

// NewAPI configures and returns application API.
func NewAPI(db *sql.DB) *API {
	userDatastore := &model.UserDatastore{DB: db}
	requestDatastore := &model.RequestDatastore{DB: db}
	itemDatastore := &model.ItemDatastore{DB: db}
	stockDataStore := &model.StockDataStore{DB: db}

	authResource := &AuthResource{UserDatastore: userDatastore}
	stockResource := &StockResource{StockDataStore: stockDataStore}
	requestResource := &RequestResource{
		requestDatastore: requestDatastore,
		userDatastore:    userDatastore,
	}
	itemResource := &ItemResource{
		ItemDatastore: itemDatastore,
		UserDatastore: userDatastore,
	}

	api := &API{
		authResource:    authResource,
		stockResource:   stockResource,
		requestResource: requestResource,
		itemResource: itemResource,
	}
	return api
}

// Router provides application routes.
func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/auth", api.authResource.router())
	r.Mount("/stocks", api.stockResource.router())
	r.Mount("/requests", api.requestResource.router())
	r.Mount("/items", api.itemResource.router())

	return r
}

type contextKey struct {
	name string
}
