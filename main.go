package main

import (
	"log"
	"net/http"

	Fast "github.com/Cthulhu-tech/fast/tree/interface/fast"
)

func main() {
	s := Fast.Create()

	if err := http.ListenAndServe(":4000", s); err != nil {
		log.Fatal(err)
	}

}
