package api

import (
	"net/http"

	"github.com/danielkrainas/gobag/context"
	gmux "github.com/gorilla/mux"

	"github.com/tenderbytes/kindermud/pkg/api/v1"
)

type Mux interface {
	http.Handler
}

type mux struct {
	router *gmux.Router
}

func NewMux() (Mux, error) {
	api := &mux{
		router: v1.RouterWithPrefix(""),
	}

	api.register(v1.RouteNameBase, http.HandlerFunc(baseHandler))
	return api, nil
}

func (api *mux) register(routeName string, dispatch http.Handler) {
	api.router.GetRoute(routeName).Handler(api.dispatcher(dispatch))
}

func (api *mux) dispatcher(dispatch http.Handler) http.Handler {
	// NOTE: a spot to add logic and decorate the route's http handler

	return dispatch
}

func (api *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOTE: a spot to add logic before and after the request is routed by the API
	defer r.Body.Close()

	w.Header().Add(v1.VersionHeader.Name, bagcontext.GetVersion(r.Context()))
	api.router.ServeHTTP(w, r)
}
