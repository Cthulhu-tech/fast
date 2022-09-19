package Fast

/*func create encapsulate method*/
func createHandler(f IFast) IServer {
	return f.create()
}

/*func create Fast handler*/
func Create() (IServer) {
	return createHandler(&Fast_{})
}

/*encapsulate method / create Fast*/
func (f *Fast_) create() IServer {
	/*logic method*/
	return &Server_{}
}
