package Fast

import "net/http"

func (f *Types) Func(callback func(w http.ResponseWriter, r *http.Request) error) {

	route[_name] = mapRouterData{method: route[_name].method, route: callback}
}

func (f *Handle) Func(callback func(w http.ResponseWriter, r *http.Request) error) {

	route[_name] = mapRouterData{method: route[_name].method, route: callback}
}
