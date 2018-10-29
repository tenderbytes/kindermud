package server

import (
	"context"
	"net/http"

	"github.com/danielkrainas/gobag/context"
	"github.com/danielkrainas/gobag/errcode"
	"github.com/rs/cors"
	"github.com/tenderbytes/kindermud/pkg/api/v1"
	"github.com/urfave/negroni"
)

func aliveHandler(path string) negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.URL.Path == path {
			w.Header().Set("Cache-Control", "no-cache")
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	})
}

func contextHandler(parent context.Context) negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		ctx := bagcontext.DefaultContextManager.Context(parent, w, r)
		defer bagcontext.DefaultContextManager.Release(ctx)

		ctx = bagcontext.WithVars(ctx, r)
		ctx = bagcontext.WithLogger(ctx, bagcontext.GetLogger(ctx))
		ctx = context.WithValue(ctx, "url.builder", v1.NewURLBuilderFromRequest(r, false))
		if iw, err := bagcontext.GetResponseWriter(ctx); err != nil {
			bagcontext.GetLogger(ctx).Warnf("response writer not found in context")
		} else {
			w = iw
		}

		next(w, r.WithContext(ctx))
	})
}

func loggingHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()
	bagcontext.GetRequestLogger(ctx).Info("request started")
	defer func() {
		status, ok := ctx.Value("http.response.status").(int)
		if ok && status >= 200 && status <= 399 {
			bagcontext.GetResponseLogger(ctx).Infof("response completed")
		}
	}()

	next(w, r)
}

func corsHandler(debug bool, config *CORSConfig) negroni.Handler {
	handler := cors.New(cors.Options{
		AllowedOrigins:   config.Origins,
		AllowedMethods:   config.Methods,
		AllowCredentials: true,
		AllowedHeaders:   config.Headers,
		Debug:            debug,
	})

	return handler
}

func trackErrorsHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := bagcontext.ErrorTracking(r.Context())
	next(w, r.WithContext(ctx))
	if errors := bagcontext.GetErrors(ctx); errors.Len() > 0 {
		if err := errcode.ServeJSON(w, errors); err != nil {
			bagcontext.GetLogger(ctx).Errorf("error serving error json: %v (from %s)", err, errors)
		}

		logErrors(ctx, errors)
	}
}

func logErrors(ctx context.Context, errors errcode.Errors) {
	for _, err := range errors {
		var lctx context.Context

		switch err.(type) {
		case errcode.Error:
			e, _ := err.(errcode.Error)
			lctx = bagcontext.WithValue(ctx, "err.code", e.Code)
			lctx = bagcontext.WithValue(lctx, "err.message", e.Code.Message())
			lctx = bagcontext.WithValue(lctx, "err.detail", e.Detail)
		case errcode.ErrorCode:
			e, _ := err.(errcode.ErrorCode)
			lctx = bagcontext.WithValue(ctx, "err.code", e)
			lctx = bagcontext.WithValue(lctx, "err.message", e.Message())
		default:
			// normal "error"
			lctx = bagcontext.WithValue(ctx, "err.code", errcode.ErrorCodeUnknown)
			lctx = bagcontext.WithValue(lctx, "err.message", err.Error())
		}

		lctx = bagcontext.WithLogger(ctx, bagcontext.GetLogger(lctx,
			"err.code",
			"err.message",
			"err.detail"))

		bagcontext.GetResponseLogger(lctx).Errorf("response completed with error")
	}
}
