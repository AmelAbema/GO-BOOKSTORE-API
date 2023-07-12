package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err1 := io.ReadAll(r.Body); err1 != nil {
		if err2 := json.Unmarshal(body, x); err2 != nil {
			return
		}
	}
}
