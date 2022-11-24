package usrdelivery

import (
	"eventool/internal/pkg/domain"
	
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

func (handler *UserHandler) GetBasicInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
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
