package eventdelivery

const (
	eventUrl       = "/event"
	CreateEventUrl = eventUrl + "/create"
	GetEventUrl    = eventUrl + "/get/{category:[a-z]+}"
	GetCatagoryUrl    = eventUrl + "/category"
	// DeleteEventUrl = eventUrl + "/delete"
	// AlterEventUrl  = eventUrl + "/alter"
)
