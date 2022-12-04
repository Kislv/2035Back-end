package usrusecase

import (
	"eventool/internal/pkg/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func InitUsrUsc(u domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepo: u,
	}
}

func (uc UserUsecase) GetBasicInfo(id uint64) (domain.User, error) {
	us, err := uc.userRepo.GetById(id)
	if err != nil {
		return domain.User{}, domain.Err.ErrObj.InternalServer
	}
	return us.ClearPasswords(), nil
}

func (uc UserUsecase) GetUserInfo(id uint64) (domain.User, error) {
	userInfo, err := uc.GetBasicInfo(id)
	if err != nil {
		return domain.User{}, err
	}

	userInfo.Categories, err = uc.userRepo.GetCategory(id)
	if err != nil {
		return domain.User{}, err
	}

	return userInfo, nil
}

func (uc UserUsecase) UpdateAvatar(clientID uint64, url string) (domain.User, error) {
	us, err := uc.userRepo.UpdateAvatar(clientID, url)
	if err != nil {
		return domain.User{}, err
	}
	return us, nil
}
