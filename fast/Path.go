package Fast

import Utils "github.com/Cthulhu-tech/fast/tree/interface/fast/utils"

/*Global method*/
func (f *F) Path(path string) IFastUniversal {

	formatingPath := Utils.CreatePath(path)

	route[formatingPath] = RouteData{}

	return f
}
