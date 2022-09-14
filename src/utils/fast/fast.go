package Fast

import (
	"fmt"
	"net/http"
)

var _name = ""
var route = make(map[string]mapRouterData)

func (f *Fast) Path(name string) *Types {

	_name = name

	route[_name] = mapRouterData{}

	return &Types{}
}

func (r *Types) Type(method string) *Types {

	route[_name] = mapRouterData{method: method}

	return r
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

func normalize(url string) (string, error) {

	var rangeStr = ""
	var delimeter = 0

	for _, char := range url {

		if char == '/' {
			delimeter++
			rangeStr = ""
		} else {
			rangeStr += string(char)
		}

	}

	fmt.Printf(rangeStr)

	if delimeter == 1 && len(rangeStr) == 0 {

		return "/", nil
	}

	return "", nil
}

func (f Fast) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func errorHandler(w http.ResponseWriter) { fmt.Fprintf(w, "not found") }

func Create() *Fast { return &Fast{} }
