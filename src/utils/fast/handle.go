package Fast

import (
	"net/http"
)

func handleMethod(w http.ResponseWriter, r *http.Request, name string, data mapRouterData) {

	if _, ok := data.method[r.Method]; ok {

		err := data.route(w, r)

		if err != nil {
			errorHandler(w)
		}

	} else {

		errorHandler(w)
	}

}
