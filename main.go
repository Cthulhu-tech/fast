package main

import (
	"log"
	"net/http"
)

type Fast struct{}
type Types struct{}
type Handle struct{}

type Handler interface {
	Path(name string) Types
	Methods
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Methods interface {
	Type(name string) Types
}

type Middleware interface {
	Func
	Use(struct{}) Middleware
}

type Func interface {
	Func(w http.ResponseWriter, r *http.Request)
}

type mapRouterData struct {
	method string
	route  string
}

var route = make(map[string]mapRouterData)

func (f *Fast) Path(name string) *Types {

	return &Types{}
}

func (r *Types) Type(name string) *Types {

	return r
}

func (f *Types) Use(callback func(next func())) *Handle {

	return &Handle{}
}

func (f *Handle) Use(callback func(next func())) *Handle {

	return &Handle{}
}

func (f *Handle) Func(callback func(w http.ResponseWriter, r *http.Request)) {

}

func (f Fast) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main() {

	h := Fast{}

	if err := http.ListenAndServe(":3000", h); err != nil {
		log.Fatal(err)
	}

}
