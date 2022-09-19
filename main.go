package main

import (
	"log"
	"net/http"

	Fast "github.com/Cthulhu-tech/fast/tree/interface/fast"
)

func main() {

	f := Fast.Create()
	
	f.Path("/")
	
	f.Path("/fast")
  
	if err := http.ListenAndServe(":4000", s); err != nil {
		log.Fatal(err)
	}

}
