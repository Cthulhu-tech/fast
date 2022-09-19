package Fast

import (
	"fmt"
	"net/http"
)

func (s *Server_) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "work")
}
