package routes

import (
	"esefexapi/timer"
	"esefexapi/types"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// api/playsound/<sound_id>
func (h *RouteHandlers) PostPlaySound(w http.ResponseWriter, r *http.Request, userID types.UserID) {
	// log.Printf("got /playsound request\n")
	timer.SetStart()

	vars := mux.Vars(r)
	sound_id := types.SoundID(vars["sound_id"])

	err := h.a.PlaySound(sound_id, userID)
	if err != nil {
		errorMsg := fmt.Sprintf("Error playing sound: \n%+v", err)

		log.Println(errorMsg)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	io.WriteString(w, "Play sound!\n")
	timer.MessageElapsed("Played sound")
}
