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
	// trimTitle(&eventData.Title)

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

	playlistResponse, err := eu.eventRepo.CreateEvent(eventData)
	if err != nil {
		return domain.EventCreatingResponse{}, err
	}

	return playlistResponse, nil
}

func (eu EventUsecase) GetEvent(categoryName string) (domain.EventListResponse, error) {

	feed, err := eu.eventRepo.GetEvent(categoryName)
	
	if err != nil {
		return domain.EventListResponse{}, err
	}

	return feed, nil
}
