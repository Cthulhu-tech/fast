package Fast

import (
	"fmt"
	"net/http"
)

var (
	url        = ""
	route      = make(map[string]RouteData)
	allMethods = map[string]bool{"GET": true, "POST": true, "HEAD": true, "OPTIONS": true, "PUT": true, "DELETE": true, "CONNECT": true, "TRACE": true, "PATCH": true}
)

/*encapsulate method / create Fast*/
func (f *F) create() IPath {
	return f
}

func (f *F) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "work")
}

/*func create encapsulate method*/
func createHandler(f IFast) IPath {
	return f.create()
}

/*func create Fast handler*/
func Create() IPath {
	return createHandler(&F{})
}
