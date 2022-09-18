package Fast

import (
	"fmt"
	"net/http"
)

func (s *S) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "work")
}
