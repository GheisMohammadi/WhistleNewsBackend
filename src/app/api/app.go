package api

import (
	"fmt"
	"net/http"
)

/**
* POST /api/v1/skill
 */
func (api *API) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get index")
}
