package usrdelivery

import (
	"eventool/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func SetUsrHandlers(router *mux.Router, uc domain.UserUsecase) {
	handler := &UserHandler{
		uc,
	}

	router.HandleFunc(getInfoUrl, handler.GetBasicInfo).Methods("GET", "OPTIONS")
	router.HandleFunc(avatarUrl, handler.UploadAvatar).Methods("POST", "OPTIONS")
}
