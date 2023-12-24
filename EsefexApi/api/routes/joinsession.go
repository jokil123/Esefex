package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// joinsession/<server_id>
func (routes *RouteHandlers) JoinSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	server_id := vars["server_id"]

	redirectUrl := fmt.Sprintf("%s://joinsession/%s", routes.cProto, server_id)
	response := fmt.Sprintf(`<meta http-equiv="refresh" content="0; URL=%s" />`, redirectUrl)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, response)

	log.Printf("got /joinsession request\n")
}
