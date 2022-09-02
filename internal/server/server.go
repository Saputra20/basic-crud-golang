package server

import (
	"basic-crud/internal/routers"

	productrtr "basic-crud/internal/routers/product"
	rolertr "basic-crud/internal/routers/role"
	userrtr "basic-crud/internal/routers/user"
	welcomertr "basic-crud/internal/routers/welcome"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	RegisterRouters(&r.RouterGroup)

	s := &Server{
		Engine: r,
	}

	return s
}

func RegisterRouters(g *gin.RouterGroup) {

	rg := &routers.RouterGroupRoot{
		API: g.Group("/api"),
	}

	welcomeRouteReg := &welcomertr.RouterGroupRegisterer{
		PathStr:     "/welcome",
		RouterGroup: rg,
	}

	productRouteReg := &productrtr.RouterGroupRegisterer{
		PathStr:     "/products",
		RouterGroup: rg,
	}

	roleRouteReg := &rolertr.RouterGroupRegisterer{
		PathStr:     "/roles",
		RouterGroup: rg,
	}

	userRouteReg := &userrtr.RouterGroupRegisterer{
		PathStr:     "/users",
		RouterGroup: rg,
	}

	rg.RegisterRouters(
		welcomeRouteReg,
		productRouteReg,
		roleRouteReg,
		userRouteReg,
	)
}
