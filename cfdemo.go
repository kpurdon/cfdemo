package main

import (
	"log"
	"net/http"

	"github.com/kpurdon/apiresponse"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		responder := apiresponse.NewResponder(w)
		responder.OK(nil)
	})
	log.Fatal(http.ListenAndServe(":3333", nil))
}
