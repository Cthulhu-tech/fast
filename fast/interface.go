package Fast

import "net/http"

type Callback func(w http.ResponseWriter, r *http.Request)

type Path_ struct{}
type Fast_ struct{}
type Method_ struct{}
type Server_ struct{}

type IPath interface {
	Path(path string) IPath
}

type IServer interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type IMethod interface {
	Method(method string) IMethod
}

type IFast interface {
	create() IServer
	path() IPath
	method() IMethod
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
