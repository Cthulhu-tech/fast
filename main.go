package main

import (
	"fmt"
	"log"
	"net/http"

	Fast "github.com/Cthulhu-tech/sinhronize/src/utils/fast"
)

func main() {

	f := Fast.Create()

	f.Path("/hello").Type("get").Func(func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintf(w, "/hello route")
		return nil
	})

	f.Path("/hello/hello").Type("get").Func(func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintf(w, "/hello/hello route")
		return nil
	})

	if err := http.ListenAndServe(":4000", f); err != nil {
		log.Fatal(err)
	}

}
