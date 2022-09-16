package Fast

import "net/http"

/*implements create / Path*/
type F struct {}

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
	Use(callback func()) IFastUniversal
	/*URL method*/
	Type(method string) IFastUniversal
	/*Last function*/
	Handler(func(w http.ResponseWriter, r *http.Request))
}
