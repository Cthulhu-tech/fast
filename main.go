package main

import (
	"log"
	"net/http"

	Fast "github.com/Cthulhu-tech/fast/tree/interface/fast"
)

func main() {

	f := Fast.Create()
	
	if err := http.ListenAndServe(":4000", f); err != nil {
		log.Fatal(err)
	}

}
