package Fast

import (
	"fmt"
	"net/http"
	"unicode"
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

func appendData(rangeStr string) string {

	var str bool
	var num bool
	var convert string

	if rangeStr != "" {

		for _, char := range rangeStr {
			if unicode.IsNumber(char) {
				num = true
			} else {
				str = true
			}
		}

		if num && str {
			convert = "[si]"
			return convert
		}

		if num {
			convert = "[i]"
		}

		if str {
			convert = "[s]"
		}

	}

	return convert
}

func normalize(url string) (string, error) {

	var lenUrl = len(url)
	var rangeStr string
	var delimeter = 0
	var str string

	for pos, char := range url {

		if char == '/' {

			delimeter++
			str += appendData(rangeStr)

			rangeStr = ""
		} else {
			rangeStr += string(char)
		}

		if pos == lenUrl-1 && char != '/' {

			str += appendData(rangeStr)
		}
	}

	if delimeter == 1 && rangeStr == "" {
		return "/", nil
	}

	fmt.Println(str)

	return "", nil
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
