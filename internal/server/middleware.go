package server

import (
	"GobbleGame/internal/vars"
	"mime"
	"net/http"
)

func CorrectRequest(next http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				Error(w, r, http.StatusBadRequest, vars.ErrMalformedContentType)
				return
			}
			if mt != "application/json" {
				Error(w, r, http.StatusUnsupportedMediaType, vars.ErrUnsupportedMediaType)
				return
			}
		}
		if r.Method != method {
			Error(w, r, http.StatusMethodNotAllowed, vars.ErrMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	}
}
