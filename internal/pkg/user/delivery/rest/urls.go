package usrdelivery

const (
	userUrl 	  = "/user"
	getInfoUrl    = userUrl + "/{id:[0-9]+}"
	// bookmarksUrl  = userUrl + "/bookmarks/{id:[0-9]+}"
	// updateUrl     = userUrl + "/update/{id:[0-9]+}"
	// reviewsUrl    = userUrl + "/reviews/{id:[0-9]+}"
	avatarUrl     = userUrl + "/update/avatar/{id:[0-9]+}"
	// subscribePush = userUrl + "/subscribePush"
)
