package Fast

import (
	"fmt"
	"net/http"
)

func (f *F) NewCoolMethod() {
	println("Yes i'm cool method!")
}

/*encapsulate method / create Fast*/
func (f *F) create() IPath {

	/*logic method*/
	println("working with create")

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
