package wapp

import (
	"fmt"
	"io"
	"net/http"
)

type Wapp struct {
	lastfoo []byte
}

//http.Handler interface:
func (wa *Wapp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(wa.lastfoo)

	bs, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	wa.lastfoo = bs
	fmt.Println(string(bs))
}
