package Fast

var middleware = false

func (f *Types) Use(callback func(next func())) *Handle {

	return &Handle{}
}
