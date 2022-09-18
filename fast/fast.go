package Fast

import (
	"fmt"
	"net/http"
)

/*encapsulate method / create Fast*/
func (f *F) create() IPath {

	/*logic method*/
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

	print(
		",------.               ,--.  \n" +
			"|  .---',--,--. ,---.,-'  '-.\n" +
			"|  `--,' ,-.  |(  .-''-.  .-'\n" +
			"|  |`  \\ '-'  |.-'  `) |  | \n" +
			"`--'    `--`--'`----'  `--'  \n" +
			"_____________________________\n" +
			"\n" + "Starting with: \n")
	return createHandler(&F{})
}
