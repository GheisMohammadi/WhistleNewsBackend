package api

import (
	"fmt"
	"net/http"
)

/**
* POST /v1/counter
 */
func (api *API) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, http.StatusOK)
}
