package controller

import (
	"go-gin-app/entity"
	"go-gin-app/service"
	"go-gin-app/validators"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func NewController(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

func (c *controller) FindAll() []entity.Video {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	// err := ctx.BindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
