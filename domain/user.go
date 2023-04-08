package domain

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Hash     string `json:"hash"`
	Token    string `json:"token"`
	Id       int    `json:"id"`
}

func (u *User) IsValid() bool {
	return u.Login != "" && u.Password != ""
}

// func (u *User) FromReq(req gophserver.RegisterLoginReq) {
// 	u.Login = req.Login
// 	u.Password = req.Password
// }

// type Bearer struct {
// 	Bearer string `json:"bearer"`
// }
