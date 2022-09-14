package main

import (
	"log"
	"net/http"

	Fast "github.com/Cthulhu-tech/sinhronize/src/utils/fast"
)

func main() {

	f := Fast.Create()

	f.Path("/").Type("POST").Type("GET").Use(func() {}).Func(func(w http.ResponseWriter, r *http.Request) error {
		println("/")
		return nil
	})

	f.Path("/hello").Type("POST").Type("GET").Func(func(w http.ResponseWriter, r *http.Request) error {
		println("/hello")
		return nil
	})

	if err := http.ListenAndServe(":4000", f); err != nil {
		log.Fatal(err)
	}

}
