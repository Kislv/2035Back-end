package eventdelivery

const (
	eventUrl      		  = "/event"
	CreateEventUrl		  = eventUrl + "/create"
	GetEventUrl   		  = eventUrl + "/get"
	GetCertainEventUrl    = GetEventUrl + "/{id:[0-9]+}"
	GetCatagoryUrl    = eventUrl + "/category"
	EventSignUpUrl		  = eventUrl + "/signup/{id:[0-9]+}"
	// DeleteEventUrl = eventUrl + "/delete"
	// AlterEventUrl  = eventUrl + "/alter"
)
