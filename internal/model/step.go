package model

type CreateStepReq struct {
	Id          int    `json:"id"`
	Description string `json:"description" binding:"required"`
	Duration    int64  `json:"duration" binding:"required"`
}
