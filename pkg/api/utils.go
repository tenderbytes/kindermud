package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	gobagcontext "github.com/danielkrainas/gobag/context"
)

func setError(ctx context.Context, err error) {
	gobagcontext.TrackError(ctx, err)
}

func setErrorAndWarn(ctx context.Context, err error, msg string) {
	gobagcontext.GetLogger(ctx).Warn(msg)
	setError(ctx, err)
}

func setErrorAndWarnf(ctx context.Context, err error, format string, args ...interface{}) {
	gobagcontext.GetLogger(ctx).Warnf(format, args...)
	setError(ctx, err)
}

func setErrorAndLog(ctx context.Context, err error, msg string) {
	gobagcontext.GetLogger(ctx).Error(msg)
	setError(ctx, err)
}

func setErrorAndLogf(ctx context.Context, err error, format string, args ...interface{}) {
	gobagcontext.GetLogger(ctx).Errorf(format, args...)
	setError(ctx, err)
}

func readRequestOptions(r *http.Request, options interface{}) error {
	err := json.NewDecoder(r.Body).Decode(options)
	if err != nil {
		gobagcontext.GetLogger(r.Context()).Errorf("error reading body: %v", err)
		return err
	}

	return nil
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	const emptyJSON = "{}"

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprint(len(emptyJSON)))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, emptyJSON)
}
