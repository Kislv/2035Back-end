package plausecase

import (
	"eventool/internal/pkg/domain"
	"eventool/internal/pkg/utils/cast"
	"eventool/internal/pkg/utils/log"

	// usrusecase "eventool/internal/pkg/user/usecase"
	// usrdelivery "eventool/internal/pkg/user/delivery/rest"
	// "usrdelivery"
	"strings"
)

type EventUsecase struct {
	eventRepo domain.EventRepository
}

func trimTitle(title *string) {
	*title = strings.Trim(*title, " ")
}

func InitEventUsc(pr domain.EventRepository) domain.EventUsecase {
	return &EventUsecase{
		eventRepo: pr,
	}
}

func (eu EventUsecase) CreateEvent(eventData domain.EventCreatingRequest) (domain.EventCreatingResponse, error) {
	alreadyExist, err := eu.eventRepo.EventAlreadyExist(eventData)
	if err != nil {
		return domain.EventCreatingResponse{}, err
	}

	if alreadyExist {
		return domain.EventCreatingResponse{}, domain.Err.ErrObj.PlaylistExist
	}

	if !eventData.IsValid() {
		return domain.EventCreatingResponse{}, domain.Err.ErrObj.InvalidTitle
	}

	eventCreatingResponse, err := eu.eventRepo.CreateEvent(eventData)
	if err != nil {
		return domain.EventCreatingResponse{}, err
	}

	eventCreatingResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	if err != nil {
		return domain.EventCreatingResponse{}, err
	}
	return eventCreatingResponse, nil
}

func (eu EventUsecase) GetEvent(categoriesName []string) (domain.EventListResponse, error) {
	
	feed, err := eu.eventRepo.GetEvent(categoriesName)
	
	if err != nil {
		return domain.EventListResponse{}, err
	}

	return feed, nil
}

func (eu EventUsecase) GetCertainEvent(eventId uint64) (domain.EventCreatingResponse, error) {

	// userAge, err :=  eu.eventRepo.GetUserAge(userId)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }

	// if !isValidUser {
	// 	return domain.EventCreatingResponse{}, nil
	// }

	event, err := eu.eventRepo.GetCertainEvent(eventId)
	
	if err != nil {
		return domain.EventCreatingResponse{}, err
	}

	return event, nil
}

func (eu EventUsecase) GetCategory() (domain.CategoryListResponse, error) {

	categoryList, err := eu.eventRepo.GetCategory()
	
	if err != nil {
		return domain.CategoryListResponse{}, err
	}

	return categoryList, nil
}

func (eu EventUsecase) EventSignUp(eventId uint64, userId uint64) (error)  {

	// log.Info("EventSignUp usecase: start")
	// var UserHandler *usrdelivery.UserHandler
	// userInfo, err := UserHandler.UserUsecase.GetBasicInfo(userId)
	
	
	// if err != nil {
	// 	log.Error(err)
	// 	return err
	// }
	// log.Info("EventSignUp usecase: after get basic Info." + userInfo.Username)

	event, err := eu.GetCertainEvent(eventId)
	if err != nil {
		return err
	}
	log.Info("EventSignUp usecase: after get certain event. Title = " + event.Title)


	// isValidUser, err := eu.IsUserValidForEvent(event, userInfo)
	// if (!isValidUser){
	// 	return domain.Err.ErrObj.BadInput
	// }

	log.Info("EventSignUp usecase: after isValidUser = true")

	err = eu.eventRepo.SignUpUserForEvent(eventId, userId)
	
	if err != nil {
		log.Info("EventSignUp usecase: UserAlreadySignUpForThisEvent")

		return domain.Err.ErrObj.UserAlreadySignUpForThisEvent
	}

	return nil
}

func (eu EventUsecase) IsUserValidForEvent(event domain.EventCreatingResponse, user domain.User) (bool, error)  {

	if (cast.IntToStr(user.Age) < event.MinAge){
		return false, nil
	}
	if (cast.IntToStr(user.Age) > event.MaxAge && event.MaxAge != "0"){
		return false, nil
	}

	return true, nil
}

func (eu EventUsecase) CancelEventSignUp(eventId uint64, userId uint64) (error)  {

	err := eu.eventRepo.CancelEventSignUp(eventId, userId)
	
	if err != nil {
		return domain.Err.ErrObj.UserDontSignUpForThisEvent
	}

	return nil
}
