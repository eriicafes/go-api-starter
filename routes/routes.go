package routes

import "github.com/gin-gonic/gin"

type Controller interface {
	Register(r *gin.RouterGroup)
}

type registry struct {
	path        string
	controller  Controller
	middlewares []gin.HandlerFunc
}

func Register(path string, controller Controller, middlewares ...gin.HandlerFunc) registry {
	return registry{
		path:        path,
		controller:  controller,
		middlewares: middlewares,
	}
}

func RegisterControllers(router *gin.Engine, controllers ...registry) {
	for _, registry := range controllers {
		if registry.controller == nil {
			continue
		}
		registry.controller.Register(router.Group(
			registry.path,
			registry.middlewares...,
		))
	}
}

func RegisterSubControllers(router *gin.RouterGroup, controllers ...registry) {
	for _, registry := range controllers {
		if registry.controller == nil {
			continue
		}
		registry.controller.Register(router.Group(
			registry.path,
			registry.middlewares...,
		))
	}
}
