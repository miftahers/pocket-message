package dto

type Login struct {
	Username string `json:"username" form:"username"`
	Token    string `json:"token" form:"token"`
}
