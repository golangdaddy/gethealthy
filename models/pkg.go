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
