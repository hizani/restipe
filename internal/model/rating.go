package model

type RateReq struct {
	Rating int `json:"rating" binding:"required"`
}
