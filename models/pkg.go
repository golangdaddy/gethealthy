package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirchorg/go/cloudfunc"
)

func getTime() int64 {
	return time.Now().UTC().Unix()
}

func AssertKeyValue(w http.ResponseWriter, m map[string]interface{}, key string) (string, bool) {
	s, ok := m[key].(string)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return s, false
	}
	return s, true
}

func AssertKeyValueFloat(w http.ResponseWriter, m map[string]interface{}, key string) (float64, bool) {
	f, ok := m[key].(float64)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return 0, false
	}
	return f, true
}

func AssertKeyValueBool(w http.ResponseWriter, m map[string]interface{}, key string) (bool, bool) {
	v, ok := m[key].(bool)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false, false
	}
	return v, true
}

func AssertKeyValueInt(w http.ResponseWriter, m map[string]interface{}, key string) (int, bool) {
	v, ok := m[key].(int)
	if !ok {
		err := fmt.Errorf("'%s' is required for this request", key)
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return 0, false
	}
	return v, true
}
