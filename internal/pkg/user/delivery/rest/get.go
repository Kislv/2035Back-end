package usrdelivery

import (
	"eventool/internal/pkg/domain"
	"eventool/internal/pkg/sessions"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

func (handler *UserHandler) GetBasicInfo (w http.ResponseWriter, r *http.Request) {
	userSessionId, err := sessions.CheckSession(r)
	if err == domain.Err.ErrObj.UserNotLoggedIn {

		out, err := easyjson.Marshal(domain.User{})
		if err != nil {
			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusForbidden)
		w.Write(out)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	if (userId != userSessionId) {
		out, err := easyjson.Marshal(domain.User{})
		if err != nil {
			http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusForbidden)
		w.Write(out)
		return
	}

	us, err := handler.UserUsecase.GetBasicInfo(userId)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusNotFound)
		return
	}

	out, err := easyjson.Marshal(us)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
