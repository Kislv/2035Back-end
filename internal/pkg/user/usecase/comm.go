package usrusecase

import (
	// autusecase "eventool/internal/pkg/authorization/usecase"
	"eventool/internal/pkg/domain"
	"eventool/internal/pkg/utils/cast"
	"eventool/internal/pkg/utils/log"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func InitUsrUsc(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (uc userUsecase) GetBasicInfo(id uint64) (domain.User, error) {
	log.Info("GetBasicInfo id = " + cast.IntToStr(id))
	us, err := uc.userRepo.GetById(id)
	if err != nil {
		return domain.User{}, domain.Err.ErrObj.InternalServer
	}
	log.Info("GetBasicInfo before return" + cast.IntToStr(id))
	return us.ClearPasswords(), nil
}

func (uc userUsecase) UpdateAvatar(clientID uint64, url string) (domain.User, error) {
	us, err := uc.userRepo.UpdateAvatar(clientID, url)
	if err != nil {
		return domain.User{}, err
	}
	return us, nil
}
