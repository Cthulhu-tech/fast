package Fast

import (
	"strings"
)

func (r *Types) Type(method string) *Types {

	var methodUpper = strings.ToUpper(method)

	if _, ok := allMethods[methodUpper]; !ok {

		errorPrint("method: \033[33m[ " + method + " ]\033[0m - not found")
		return r
	}

	if route[_name].method == nil {
		route[_name] = mapRouterData{method: make(map[string]bool)}
	}

	if _, ok := route[_name].method[methodUpper]; !ok {
		route[_name].method[methodUpper] = true
	}

	return r
}
