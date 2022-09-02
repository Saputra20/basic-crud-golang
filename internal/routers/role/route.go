package rolertr

import "basic-crud/internal/routers"

type RouterGroupRegisterer struct {
	PathStr     string
	RouterGroup *routers.RouterGroupRoot
}

var roleHandler = RoleHandler{}

func (rr *RouterGroupRegisterer) APIRoutes() {
	rg := rr.RouterGroup.API.Group(rr.PathStr)

	rg.GET("/", roleHandler.getRoles)
	rg.GET("/:id", roleHandler.getRole)
	rg.POST("/", roleHandler.createRole)
	rg.PUT("/:id", roleHandler.updateRole)
	rg.DELETE("/:id", roleHandler.deleteRole)
}

func (rr *RouterGroupRegisterer) Path() string {
	return rr.PathStr
}
