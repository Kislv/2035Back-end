package eventdelivery

const (
	eventUrl      		  = "/event"
	CreateEventUrl		  = eventUrl + "/create"
	GetEventUrl   		  = eventUrl + "/get"
	GetCertainEventUrl    = GetEventUrl + "/{id:[0-9]+}"
	GetRecomendedEvent    = GetEventUrl + "/recomended"
	GetCatagoryUrl        = eventUrl + "/category"
	EventSignUpUrl		  = eventUrl + "/signup/{id:[0-9]+}"
	CancelEventSignUpUrl  = eventUrl + "/signup/cancel/{id:[0-9]+}"
	// DeleteEventUrl = eventUrl + "/delete"
	// AlterEventUrl  = eventUrl + "/alter"
)
