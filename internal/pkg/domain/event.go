package domain

// TODO write valid path
const BaseEventPicture = "/home/ubuntu/lolkek/static/event/event.png"

type EventCreatingRequest struct {
	// TODO: define, what is required fields
	PosterPath             string   `json:"posterpath"`
	Title                  string   `json:"title"`
	Rating                 string   `json:"rating"`
	VotesNum               uint64   `json:"-"`
	Description            string   `json:"description"`
	UserId                 string   `json:"userid"`
	Longitude              string   `json:"longitude"`
	Latitude               string   `json:"latitude"`
	CurrentMembersQuantity uint32   `json:"currentmembersquantity"`
	MaxMembersQuantity     uint32   `json:"maxmembersquantity"`
	MinMembersQuantity     uint32   `json:"minmembersquantity"`
	CreatingDate           string   `json:"creatingdate"`
	StartDate              string   `json:"startdate"`
	EndDate                string   `json:"enddate"`
	MinAge                 string   `json:"minage"`
	MaxAge                 string   `json:"maxage"`
	Price                  string   `json:"price"`
	Categories             []string `json:"categories"`
}

type EventResponse struct {
	// TODO: define, what is required fields
	Id					   uint64   `json:"id"`
	PosterPath             string   `json:"posterpath"`
	Title                  string   `json:"title"`
	Rating                 string   `json:"rating"`
	VotesNum               uint64   `json:"-"`
	Description            string   `json:"description"`
	UserId                 string   `json:"userid"`
	Longitude              string   `json:"longitude"`
	Latitude               string   `json:"latitude"`
	CurrentMembersQuantity uint32   `json:"currentmembersquantity"`
	MaxMembersQuantity     uint32   `json:"maxmembersquantity"`
	MinMembersQuantity     uint32   `json:"minmembersquantity"`
	CreatingDate           string   `json:"creatingdate"`
	StartDate              string   `json:"startdate"`
	EndDate                string   `json:"enddate"`
	MinAge                 string   `json:"minage"`
	MaxAge                 string   `json:"maxage"`
	Price                  string   `json:"price"`
	Categories             []string `json:"categories"`
}
