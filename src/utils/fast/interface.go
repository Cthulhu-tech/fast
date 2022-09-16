package Fast

import "net/http"

type Fast struct{}
type Types struct{}
type Handle struct{}


type mapRouterData struct {
	method map[string]bool
	route  func(w http.ResponseWriter, r *http.Request) error
}
