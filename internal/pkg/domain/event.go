package domain

// TODO write valid path
const (
	BaseEventPicture = "/home/ubuntu/lolkek/static/event/event.png"
	maxEventTitleLength = 200
)

type EventCreatingRequest struct {
	// TODO: define, what is required fields
	Title                  string   `json:"title"`
	Description            string   `json:"description,omitempty"`
	UserId                 string   `json:"userid"`
	Longitude              string   `json:"longitude"`
	Latitude               string   `json:"latitude"`
	MaxMembersQuantity     uint32   `json:"maxmembersquantity,omitempty"`
	MinMembersQuantity     uint32   `json:"minmembersquantity,omitempty"`
	StartDate              string   `json:"startdate"`
	EndDate                string   `json:"enddate,omitempty"`
	MinAge                 string   `json:"minage,omitempty"`
	MaxAge                 string   `json:"maxage,omitempty"`
	Price                  string   `json:"price,omitempty"`
	Categories             []string `json:"categories"`
}

func (er *EventCreatingRequest) SetDefault() () {
	er.Description = ""
	er.MaxMembersQuantity = 0
	er.MinMembersQuantity = 0
	er.EndDate = "1970-01-01 00:00:00"
	er.MaxAge = "0"
	er.MinAge = "0"
	er.Price = "0"

	return
}
func (er EventCreatingRequest) IsValid() (isValid bool) {
	if len(er.Title) > maxEventTitleLength {
		return false
	}
	// TODO: add checking of startdate < enddate 
	return true
}

type EventCreatingResponse struct {
	Id                     uint64   `json:"id"`
	PosterPath             string   `json:"posterpath"`
	Title                  string   `json:"title"`
	Rating                 string   `json:"rating"`
	VotesNum               uint64   `json:"-"`
	Description            string   `json:"description"`
	UserId                 string   `json:"userid"`
	Longitude              string   `json:"longitude"`
	Latitude               string   `json:"latitude"`
	CurrentMembersQuantity uint64   `json:"currentmembersquantity"`
	MaxMembersQuantity     uint64   `json:"maxmembersquantity"`
	MinMembersQuantity     uint64   `json:"minmembersquantity"`
	CreatingDate           string   `json:"creatingdate"`
	StartDate              string   `json:"startdate"`
	EndDate                string   `json:"enddate"`
	MinAge                 string   `json:"minage"`
	MaxAge                 string   `json:"maxage"`
	Price                  string   `json:"price"`
	Categories             []string `json:"categories"`
}


type EventListResponse struct {
	EventList []EventCreatingResponse `json:"eventlist"`
}

type CategoryResponse struct {
	Name             string   `json:"name"`
	ImagePath             string   `json:"ImagePath"`
}

type CategoryListResponse struct {
	CategoryList []CategoryResponse `json:"categorylist"`
}

type EventRepository interface {
	CreateEvent(event EventCreatingRequest) (EventCreatingResponse, error)
	EventAlreadyExist(event EventCreatingRequest) (bool, error)
	GetEvent(categoriesName []string) (EventListResponse, error)  
	GetCategory() (CategoryListResponse, error)
	CreateEventCategory(eventId uint64, categories []string) ([]string, error)
	GetCertainEvent(eventId uint64) (EventCreatingResponse, error)
	SignUpUserForEvent(eventId uint64, userId uint64) (error)
	// GetUserAge (userId uint64) (, error)


	// AddMovie(addMovieInfo MovieInPlaylist) error
	// DeleteMovie(movieDeleteInfo MovieInPlaylist) error
	// DeletePlaylist(deletePlaylistInfo DeletePlaylistInfo) error
	// PlaylistAlreadyExist(playlist PlaylistRequest) (bool, error)
	// AlterPlaylistPublic(alterPlaylistPublicInfo AlterPlaylistPublicInfo) error
	// AlterPlaylistTitle(alterPlaylistTitleInfo AlterPlaylistTitleInfo) error
}

type EventUsecase interface {
	CreateEvent(event EventCreatingRequest) (EventCreatingResponse, error)
	GetEvent(categoriesName []string) (EventListResponse, error) 
	GetCategory() (CategoryListResponse, error)
	GetCertainEvent(eventId uint64) (EventCreatingResponse, error)
	EventSignUp(eventId uint64, userId uint64)(error)
	
	
	
	// AddMovie(addMovieInfo MovieInPlaylist) error
	// DeleteMovie(MovieInPlaylist MovieInPlaylist) error
	// DeletePlaylist(deletePlaylistInfo DeletePlaylistInfo) error
	// AlterPlaylistPublic(alterPlaylistPublicInfo AlterPlaylistPublicInfo) error
	// AlterPlaylistTitle(alterPlaylistTitleInfo AlterPlaylistTitleInfo) error
}
