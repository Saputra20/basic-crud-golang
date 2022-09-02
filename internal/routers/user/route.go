package userrtr

import "basic-crud/internal/routers"

type RouterGroupRegisterer struct {
	PathStr     string
	RouterGroup *routers.RouterGroupRoot
}

var userHandler = UserHandler{}

func (rr *RouterGroupRegisterer) APIRoutes() {
	rg := rr.RouterGroup.API.Group(rr.PathStr)

	rg.GET("/", userHandler.getUsers)
	rg.GET("/:id", userHandler.getUser)
	rg.POST("/", userHandler.createUser)
	rg.PUT("/:id", userHandler.updateUser)
	rg.DELETE("/:id", userHandler.deleteUser)
}

func (rr *RouterGroupRegisterer) Path() string {
	return rr.PathStr
}
