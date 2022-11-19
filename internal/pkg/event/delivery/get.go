package eventdelivery

import (
	"eventool/internal/pkg/domain"
	"eventool/internal/pkg/sessions"
	"eventool/internal/pkg/utils/sanitizer"

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
	// TODO add easyjson
	err = easyjson.Unmarshal(b, EventCreatingRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
		return
	}

	sanitizer.SanitizeEventCreating(EventCreatingRequest)

	us, err := handler.EventUsecase.CreateEvent(*EventCreatingRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := easyjson.Marshal(us)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

// func (handler *EventHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
// 		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	playlistRequest := new(domain.PlaylistRequest)
// 	err = easyjson.Unmarshal(b, playlistRequest)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	sanitizer.SanitizePlaylistCreating(playlistRequest)

// 	us, err := handler.EventUsecase.CreatePlaylist(*playlistRequest)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	out, err := easyjson.Marshal(us)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(out)
// }

// func (handler *EventHandler) AddMovie(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
// 		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	b, err := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	addPlaylistInfo := new(domain.MovieInPlaylist)
// 	err = easyjson.Unmarshal(b, addPlaylistInfo)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = handler.EventUsecase.AddMovie(*addPlaylistInfo)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }

// func (handler *EventHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
// 		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	b, err := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	movieInPlaylist := new(domain.MovieInPlaylist)
// 	err = easyjson.Unmarshal(b, movieInPlaylist)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = handler.EventUsecase.DeleteMovie(*movieInPlaylist)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func (handler *EventHandler) DeletePlaylist(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
// 		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	b, err := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	deletePlaylistInfo := new(domain.DeletePlaylistInfo)
// 	err = easyjson.Unmarshal(b, deletePlaylistInfo)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = handler.EventUsecase.DeletePlaylist(*deletePlaylistInfo)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func (handler *EventHandler) AlterPlaylistPublic(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
// 		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	b, err := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	alterPlaylistPublicInfo := new(domain.AlterPlaylistPublicInfo)
// 	err = easyjson.Unmarshal(b, alterPlaylistPublicInfo)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = handler.EventUsecase.AlterPlaylistPublic(*alterPlaylistPublicInfo)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func (handler *EventHandler) AlterPlaylistTitle(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
// 		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	alterPlaylistTitleInfo := new(domain.AlterPlaylistTitleInfo)
// 	err := json.NewDecoder(r.Body).Decode(&alterPlaylistTitleInfo)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err = handler.EventUsecase.AlterPlaylistTitle(*alterPlaylistTitleInfo)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
