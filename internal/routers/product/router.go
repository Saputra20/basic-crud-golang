package productrtr

import "basic-crud/internal/routers"

type RouterGroupRegisterer struct {
	PathStr     string
	RouterGroup *routers.RouterGroupRoot
}

var productHandler = ProductHandler{}

func (rr *RouterGroupRegisterer) APIRoutes() {
	rg := rr.RouterGroup.API.Group(rr.PathStr)

	rg.GET("/", productHandler.getProducts)
	rg.GET("/:id", productHandler.getProduct)
	rg.POST("/", productHandler.createProduct)
	rg.PUT("/:id", productHandler.updateProduct)
	rg.DELETE("/:id", productHandler.deleteProduct)
}

func (rr *RouterGroupRegisterer) Path() string {
	return rr.PathStr
}
