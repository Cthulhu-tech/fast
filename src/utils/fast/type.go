package Fast

import "net/http"

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
	route  func(w http.ResponseWriter, r *http.Request) error
}
