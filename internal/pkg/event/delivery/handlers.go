package eventdelivery

import (
	"eventool/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type EventHandler struct {
	EventUsecase domain.EventUsecase
}

func SetEventHandlers(router *mux.Router, pu domain.EventUsecase) {
	handler := &EventHandler{
		EventUsecase: pu,
	}
	router.HandleFunc(CreateEventUrl, handler.CreateEvent).Methods("POST", "OPTIONS")
	router.HandleFunc(GetEventUrl, handler.GetEvent).Methods("GET", "OPTIONS")
	router.HandleFunc(GetCertainEventUrl, handler.GetCertainEvent).Methods("GET", "OPTIONS")
	router.HandleFunc(GetCatagoryUrl, handler.GetCategory).Methods("GET", "OPTIONS")
	router.HandleFunc(EventSignUpUrl, handler.EventSignUp).Methods("POST", "OPTIONS")
	router.HandleFunc(CancelEventSignUpUrl, handler.CancelEventSignUp).Methods("POST", "OPTIONS")
	
	// router.HandleFunc(DeleteEventUrl, handler.DeleteEvent).Methods("GET", "OPTIONS")
	// router.HandleFunc(AlterEventUrl, handler.AlterEvent).Methods("GET", "OPTIONS")
}
