package main

import (
	"github.com/borkab/wapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":8080", &wapp.Wapp{}) //a program a 8080 porton figyel a helyi szerveren
}
