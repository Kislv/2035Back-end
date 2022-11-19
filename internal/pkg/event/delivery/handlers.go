package eventdelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type EventHandler struct {
	EventUsecase domain.EventUsecase
}

func SetPlaHandlers(router *mux.Router, pu domain.EventUsecase) {
	handler := &EventHandler{
		EventUsecase: pu,
	}
	router.HandleFunc(CreateEventUrl, handler.CreateEvent).Methods("POST", "OPTIONS")
	router.HandleFunc(GetEventUrl, handler.GetEvent).Methods("GET", "OPTIONS")
	// router.HandleFunc(DeleteEventUrl, handler.DeleteEvent).Methods("GET", "OPTIONS")
	// router.HandleFunc(AlterEventUrl, handler.AlterEvent).Methods("GET", "OPTIONS")

	// router.HandleFunc(createPlaylistUrl, handler.CreatePlaylist).Methods("POST", "OPTIONS")
	// router.HandleFunc(addMovieUrl, handler.AddMovie).Methods("POST", "OPTIONS")
	// router.HandleFunc(deleteMovieUrl, handler.DeleteMovie).Methods("POST", "OPTIONS")
	// router.HandleFunc(deletePlaylistUrl, handler.DeletePlaylist).Methods("POST", "OPTIONS")
	// router.HandleFunc(alterPlaylistPublicUrl, handler.AlterPlaylistPublic).Methods("POST", "OPTIONS")
	// router.HandleFunc(alterPlaylistTitleUrl, handler.AlterPlaylistTitle).Methods("POST", "OPTIONS")
}
