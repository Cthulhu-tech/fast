package Fast

import "net/http"

/*implements create / Path*/
type F struct{}
type P struct{}
type S struct{}

type FuncCallback func()

type FuncCallbackHandler func(w http.ResponseWriter, r *http.Request)

type IPath interface {
	Path(path string) IFastUniversal
}

type IServer interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
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
