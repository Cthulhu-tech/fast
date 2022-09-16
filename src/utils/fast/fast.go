package Fast

import (
	"fmt"
	"net/http"
)

var _name = ""
var route = make(map[string]mapRouterData)
var allMethods = map[string]bool{"GET": true, "POST": true, "HEAD": true, "OPTIONS": true, "PUT": true, "DELETE": true, "CONNECT": true, "TRACE": true, "PATCH": true}

func (f *Fast) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if val, ok := route[r.URL.Path]; ok {

		handleMethod(w, r, r.URL.Path, val)

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
