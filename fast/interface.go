package Fast

import "net/http"

/*implements create / Path*/
type F struct{}

type FuncCallback func()

type FuncCallbackHandler func(w http.ResponseWriter, r *http.Request)

type IPath interface {
	Path(path string) IFastUniversal
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type IFast interface {
	create() IPath
	IPath
}

type IFastUniversal interface {
	/*Middleware*/
	Use(callback FuncCallback) IFastUniversal
	/*URL method*/
	Type(method string) IFastUniversal
	/*Last function*/
	Handler(calback FuncCallbackHandler)
}

type RouteData struct {
	method map[string]bool
	route  FuncCallbackHandler
}
