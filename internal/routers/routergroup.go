package routers

import "github.com/gin-gonic/gin"

type RouterGroupRegister interface {
	APIRoutes()
	Path() string
}

type RouterGroupRoot struct {
	API *gin.RouterGroup
}

func (rg *RouterGroupRoot) RegisterRouters(regs ...RouterGroupRegister) {
	for _, reg := range regs {
		reg.APIRoutes()
	}
}
