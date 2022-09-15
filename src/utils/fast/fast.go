package Fast

import (
	"fmt"
	"net/http"
)

var _name = ""
var route = make(map[string]mapRouterData)
var allMethods = map[string]bool{"GET": true, "POST": true, "HEAD": true, "OPTIONS": true, "PUT": true, "DELETE": true, "CONNECT": true, "TRACE": true, "PATCH": true}

func (f *Fast) Path(name string) *Types {

	_name = name

	route[_name] = mapRouterData{}

	return &Types{}
}

func (f *Types) Use(callback func()) *Handle {

	return &Handle{}
}

func (f *Types) Func(callback func(w http.ResponseWriter, r *http.Request) error) {

	route[_name] = mapRouterData{route: callback}

}

func (f *Handle) Use(callback func()) *Handle {

	return &Handle{}
}

func (f *Handle) Func(callback func(w http.ResponseWriter, r *http.Request) error) {

	route[_name] = mapRouterData{route: callback}

}

func (f Fast) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if val, ok := route[r.URL.Path]; ok {

		err := val.route(w, r)

		if err != nil {
			println(err)
		}

	} else {

		normalize, err := normalize(r.URL.Path)

		if err != nil {
			errorHandler(w)
		}

		if val, ok := route[normalize]; ok {

			err := val.route(w, r)

			if err != nil {
				println(err)
			}

		} else {
			errorHandler(w)
		}

	}

}

func errorHandler(w http.ResponseWriter) { fmt.Fprintf(w, "not found") }

func Create() *Fast { return &Fast{} }
