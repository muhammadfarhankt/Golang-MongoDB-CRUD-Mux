package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, employee interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), employee); err != nil {
			//fmt.Println("Error parsing the body", err)
			return
		}
	}
}
