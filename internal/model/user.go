package model

type User struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type SignupUserReq struct {
	Name     string `json:"name" binding:"required"`
	Login    string `json:"Login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SigninUserReq struct {
	Login    string `json:"Login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
