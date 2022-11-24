package eventdelivery

const (
	eventUrl      		  = "/event"
	CreateEventUrl		  = eventUrl + "/create"
	GetEventUrl   		  = eventUrl + "/get"
	GetCertainEventUrl    = GetEventUrl + "/{id:[0-9]+}"
	GetCatagoryUrl    = eventUrl + "/category"
	// DeleteEventUrl = eventUrl + "/delete"
	// AlterEventUrl  = eventUrl + "/alter"
)
