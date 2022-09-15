package Fast

var middleware = false

func (n *NextStruct) Func() {
	middleware = true
}

func (f *Handle) Use(callback func(next *Next)) *Handle {

	if middleware {
		println("work")
	}

	return &Handle{}
}

func (f *Types) Use(callback func(next func())) *Handle {

	return &Handle{}
}
