package main

import (
	"log"
	"net/http"

	Fast "github.com/Cthulhu-tech/fast/tree/interface/fast"
)

func main() {
z
	f := Fast.Create()
	
	f.Path("/")
	
	f.Path("/fast")

	if err := http.ListenAndServe(":4000", f); err != nil {
		log.Fatal(err)
	}

}
