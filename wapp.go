package wapp

import (
	"io"
	"net/http"
)

type Wapp struct {
	lastfoo []byte
}

//http.Handler interface:
func (wa *Wapp) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" { //megjegyezzuk
		bs, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		wa.lastfoo = bs
		w.WriteHeader(201)
	}
	if r.Method == "GET" { //visszaadjuk
		w.WriteHeader(http.StatusOK)
		w.Write(wa.lastfoo)
	}
}
