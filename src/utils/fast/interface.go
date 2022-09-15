package Fast

import "net/http"

type Fast struct{}
type Types struct{}
type Handle struct{}

type NextStruct struct{}

type Handler interface {
	Path(name string) Types
	Methods
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Methods interface {
	Type(name string) Types
}

type mapRouterData struct {
	method map[string]bool
	route  func(w http.ResponseWriter, r *http.Request) error
}

type Next interface {
	Func()
}
