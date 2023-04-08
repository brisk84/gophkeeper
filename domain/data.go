package domain

type Data struct {
	Id    int
	Title string
	Type  string
}

type LoginInfo struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type CardInfo struct {
	CardNo string `json:"card_no"`
	Valid  string `json:"valid"`
	CVV    string `json:"cvv"`
}
