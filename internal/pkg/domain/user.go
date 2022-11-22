package domain

// TODO write valid path
const BaseUserPicture = "/home/ubuntu/lolkek/static/avatars/profile.svg"

type User struct {
	Id             uint64   `json:"ID"`
	Username       string   `json:"username"`
	Password       string   `json:"password,omitempty"`
	RepeatPassword string   `json:"repeatpassword,omitempty"`
	Email          string   `json:"email"`
	Imgsrc         string   `json:"imgsrc"`
	PhoneNumber    string   `json:"phonenumber"`
	Categories     []string `json:"categories"`
}

func (us *User) ClearPasswords() User {
	us.Password = ""
	us.RepeatPassword = ""
	return *us
}

type UserBasic struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPublicInfo struct {
	Id       uint64 `json:"ID"`
	Username string `json:"name"`
}

type UpdUser struct {
	Username string 		`json:"username"`
	PhoneNumber    string 	`json:"phonenumber"`
	Imgsrc   string 		`json:"imgsrc"`
}

type UserNotificationToken struct {
	Token string `json:"token"`
}

type UserRepository interface {
	GetById(id uint64) (User, error)
	UpdateUser(id uint64, upd UpdUser) (User, error)
	UpdateAvatar(id uint64, url string) (User, error)
}

type UserUsecase interface {
	GetBasicInfo(id uint64) (User, error)
	UpdateUser(id uint64, upd UpdUser) (User, error)
	UpdateAvatar(id uint64, url string) (User, error)
}
