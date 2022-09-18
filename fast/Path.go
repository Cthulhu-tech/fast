package Fast

/*encapsulate method / create Fast*/
func (f *Fast_) path() IPath {
	/*logic method*/
	return &Path_{}
}

/*func create Fast Path*/
func (p *Path_) Path(path string) IPath {
	return createPath(&Fast_{})
}

/*func create encapsulate method*/
func createPath(f IFast) IPath {
	return f.path()
}
