package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Register(r *gin.RouterGroup)
}

type Record struct {
	Path        string
	Controller  Controller
	Middlewares []gin.HandlerFunc
}

func NewRecord(path string, controller Controller, middlewares ...gin.HandlerFunc) Record {
	return Record{
		Path:        path,
		Controller:  controller,
		Middlewares: middlewares,
	}
}

func Register(router *gin.Engine, records ...Record) {
	for _, record := range records {
		if record.Controller == nil {
			continue
		}
		record.Controller.Register(router.Group(
			record.Path,
			record.Middlewares...,
		))
	}
}

func RegisterSub(router *gin.RouterGroup, controllers ...Record) {
	for _, record := range controllers {
		if record.Controller == nil {
			continue
		}
		record.Controller.Register(router.Group(
			record.Path,
			record.Middlewares...,
		))
	}
}
