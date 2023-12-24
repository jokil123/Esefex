package routes

import (
	"esefexapi/sounddb"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// api/playsound/<user_id>/<server_id>/<sound_id>
func (routes *RouteHandlers) PlaySound(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /playsound request\n")

	vars := mux.Vars(r)
	user_id := vars["user_id"]
	server_id := vars["server_id"]
	sound_id := vars["sound_id"]

	err := routes.a.PlaySound(sounddb.SuidFromStrings(server_id, sound_id), server_id, user_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	io.WriteString(w, "Play sound!\n")
}