package eventdelivery

import (
	"eventool/internal/pkg/domain"
	"eventool/internal/pkg/sessions"
	"eventool/internal/pkg/utils/sanitizer"

	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"

	"github.com/mailru/easyjson"
)

func (handler *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	EventCreatingRequest := new(domain.EventCreatingRequest)
	err = easyjson.Unmarshal(b, EventCreatingRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
		return
	}

	sanitizer.SanitizeEventCreating(EventCreatingRequest)

	es, err := handler.EventUsecase.CreateEvent(*EventCreatingRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *EventHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryName := params["category"]
	
	eventList, err := handler.EventUsecase.GetEvent(categoryName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	out, err := easyjson.Marshal(eventList)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
