package app

import "net/http"

func Response(rw http.ResponseWriter, sc int) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(sc)
	return nil
}
