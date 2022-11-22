package sanitizer

import (
	"eventool/internal/pkg/domain"

	"github.com/microcosm-cc/bluemonday"
)

func SanitizeUser(user *domain.User) {
	sanitizer := bluemonday.UGCPolicy()
	user.Email = sanitizer.Sanitize(user.Email)
	user.Username = sanitizer.Sanitize(user.Username)
	user.PhoneNumber = sanitizer.Sanitize(user.PhoneNumber)
	user.Imgsrc = sanitizer.Sanitize(user.Imgsrc)
}

func SanitizeUpdUser(user *domain.UpdUser) {
	sanitizer := bluemonday.UGCPolicy()
	user.Username = sanitizer.Sanitize(user.Username)
	user.PhoneNumber = sanitizer.Sanitize(user.PhoneNumber)
	user.Imgsrc = sanitizer.Sanitize(user.Imgsrc)
}

func SanitizeComment(comm *string) {
	sanitizer := bluemonday.UGCPolicy()
	*comm = sanitizer.Sanitize(*comm)
}

func SanitizeUserBasic(login *domain.UserBasic) {
	sanitizer := bluemonday.UGCPolicy()
	login.Email = sanitizer.Sanitize(login.Email)
}

func SanitizeEventCreating(event *domain.EventCreatingRequest) {
	sanitizer := bluemonday.UGCPolicy()
	event.Title = sanitizer.Sanitize(event.Title)
	event.Description = sanitizer.Sanitize(event.Description)
	event.UserId = sanitizer.Sanitize(event.UserId)
	event.Longitude = sanitizer.Sanitize(event.Longitude)
	event.Latitude = sanitizer.Sanitize(event.Latitude)
	event.StartDate = sanitizer.Sanitize(event.StartDate)
	event.EndDate = sanitizer.Sanitize(event.EndDate)
	event.MinAge = sanitizer.Sanitize(event.MinAge)
	event.MaxAge = sanitizer.Sanitize(event.MaxAge)
	event.Price = sanitizer.Sanitize(event.Price)
	for i, _ := range(event.Categories){
		event.Categories[i] = sanitizer.Sanitize(event.Categories[i])
	}
}
