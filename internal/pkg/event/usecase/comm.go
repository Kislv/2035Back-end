package plausecase

import (
	"eventool/internal/pkg/domain"
	"eventool/internal/pkg/utils/cast"
	"eventool/internal/pkg/utils/log"

	// usrusecase "eventool/internal/pkg/user/usecase"
	// usrdelivery "eventool/internal/pkg/user/delivery/rest"
	// usrusecase "eventool/internal/pkg/user/usecase"
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

	userAge, err := eu.eventRepo.GetUserAge(userId)
	if err != nil {
		return err
	}
	eventMinAge, eventMaxAge, err := eu.eventRepo.GetEventAges(eventId)
	if err != nil {
		return err
	}


	isValidUser, err := eu.IsUserValidForEvent(eventMinAge, eventMaxAge, userAge)
	if err != nil {
		return err
	}
	if (!isValidUser){
		return domain.Err.ErrObj.BadInput
	}


	err = eu.eventRepo.SignUpUserForEvent(eventId, userId)
	if err != nil {
		return domain.Err.ErrObj.UserAlreadySignUpForThisEvent
	}

	return nil
}

func (eu EventUsecase) IsUserValidForEvent(minAge uint16, maxAge uint16, age uint64) (bool, error)  {
	if (cast.IntToStr(age) < cast.Uint16ToStr(minAge)){
		return false, nil
	}
	if (cast.IntToStr(age) > cast.Uint16ToStr(maxAge) && cast.Uint16ToStr(maxAge) != "0"){
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

func (eu EventUsecase) GetRecomendedEvent(userId uint64) (domain.EventListResponse, error) {

	categories, err := eu.eventRepo.GetUserCategory(userId)
	
	if err != nil {
		log.Error(err)
		return domain.EventListResponse{}, err
	}
	
	eventList, err := eu.GetEvent(categories)
	
	if err != nil {
		return domain.EventListResponse{}, err
	}

	return eventList, nil
}
