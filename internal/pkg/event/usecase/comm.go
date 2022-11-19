package plausecase

import (
	"codex/internal/pkg/domain"
	"strings"
)

type EventUsecase struct {
	playlistRepo domain.Plarepository
}

func trimTitle(title *string) {
	*title = strings.Trim(*title, " ")
}

func InitPlaUsc(pr domain.Plarepository) domain.EventUsecase {
	return &EventUsecase{
		playlistRepo: pr,
	}
}

func (pu EventUsecase) CreateEvent (eventData domain.EventRequest) (domain.EventResponse, error) {
	trimTitle(&eventData.Title)

	alreadyExist, err := pu.playlistRepo.PlaylistAlreadyExist(playlistData)
	if err != nil {
		return domain.PlaylistResponse{}, err
	}

	if alreadyExist {
		return domain.PlaylistResponse{}, domain.Err.ErrObj.PlaylistExist
	}

	if !playlistData.TitleIsValid() {
		return domain.PlaylistResponse{}, domain.Err.ErrObj.InvalidTitle
	}

	playlistResponse, err := pu.playlistRepo.CreatePlaylist(playlistData)
	if err != nil {
		return domain.PlaylistResponse{}, err
	}

	return playlistResponse, nil
}
