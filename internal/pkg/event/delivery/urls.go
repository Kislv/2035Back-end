package eventdelivery

const (
	eventUrl       = "/event"
	CreateEventUrl = eventUrl + "/create"
	GetEventUrl    = eventUrl + "/get/{category:[a-z]+}"
	// DeleteEventUrl = eventUrl + "/delete"
	// AlterEventUrl  = eventUrl + "/alter"
)
