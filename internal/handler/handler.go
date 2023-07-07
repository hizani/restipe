package handler

import (
	"net/http"
	"restipe/internal/handler/ginhandler"
	"restipe/internal/service"
)

func NewGinHandler(service *service.Service) http.Handler {
	return ginhandler.New(service)

}
