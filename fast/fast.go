package Fast

var (
	url        = ""
	route      = make(map[string]RouteData)
	allMethods = map[string]bool{"GET": true, "POST": true, "HEAD": true, "OPTIONS": true, "PUT": true, "DELETE": true, "CONNECT": true, "TRACE": true, "PATCH": true}
)

/*encapsulate method / create Fast*/
func (p Path) create() (IPath, IServer) {
	return p, &S{}
}

/*func create Fast handler*/
func Create() (IPath, IServer) {
	return p.create()
}
