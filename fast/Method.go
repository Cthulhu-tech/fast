package Fast

/*encapsulate method / create Fast*/
func (f *Fast_) method() IMethod {
	/*logic method*/
	return &Method_{}
}

/*func create Fast Path*/
func (p *Method_) Method(method string) IMethod {
	return createMethod(&Fast_{})
}

/*func create encapsulate method*/
func createMethod(f IFast) IMethod {
	return f.method()
}
