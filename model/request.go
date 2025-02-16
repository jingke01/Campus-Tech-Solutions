package model

type RegisterReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
