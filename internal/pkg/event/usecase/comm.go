package plausecase

import (
	"eventool/internal/pkg/domain"
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
	// event, err := handler.EventUsecase.CheckUserPermissionOnEvent(userId)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// check user permission

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
