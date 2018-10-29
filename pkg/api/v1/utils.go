package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/danielkrainas/gobag/mapconv"
)

func fieldInvalid(fieldName string, desc string) error {
	return fmt.Errorf("%q invalid: %s", fieldName, desc)
}

func fieldMissingOrInvalid(fieldName string) error {
	return fmt.Errorf("%q missing or invalid", fieldName)
}

func fieldIsNonNegative(fieldName string) error {
	return fmt.Errorf("%q must be non-negative", fieldName)
}

func parseQueryArray(vals url.Values, keyName string) []string {
	arrKey := keyName + "[]"
	if v, ok := vals[keyName]; ok {
		return v
	} else if v, ok := vals[arrKey]; ok {
		return v
	}

	return []string{}
}

type V1Mapper interface {
	V1Map() map[string]interface{}
}

func Serve(w http.ResponseWriter, data interface{}) error {
	mdata := data
	if v1mapper, ok := data.(V1Mapper); ok {
		mdata = v1mapper.V1Map()
	} else if mapper, ok := data.(mapconv.Mapper); ok {
		mdata = mapper.Map()
	}

	return ServeJSON(w, mdata)
}

func ServeJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(data)
}
