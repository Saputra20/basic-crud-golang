package welcomertr

import "basic-crud/internal/routers"

type RouterGroupRegisterer struct {
	PathStr     string
	RouterGroup *routers.RouterGroupRoot
}

func (rr *RouterGroupRegisterer) APIRoutes() {
	rg := rr.RouterGroup.API.Group(rr.PathStr)

	rg.GET("/", welcomeHandler)
}

func (rr *RouterGroupRegisterer) Path() string {
	return rr.PathStr
}
