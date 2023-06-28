package http

import (
	"github.com/gin-gonic/gin"
)

type http struct {
	Engine *gin.Engine
}

func NewHttp() *http {

	r := gin.Default()

	return &http{
		Engine: r,
	}
}

type RegisterGinFunc func(engine *gin.Engine)

func (h *http) Register(servers ...RegisterGinFunc) *http {
	for _, registerServer := range servers {
		registerServer(h.Engine)
	}
	return h
}

func (h *http) Run(addr string) {
	h.Engine.Run(addr)
}
