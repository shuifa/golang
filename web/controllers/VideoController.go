package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"githun.com/oushuifa/golang/web/models"
)

type Generator struct {
	counter int
}

func (g *Generator) GetNextId() int {
	g.counter++
	return g.counter
}

var generator = &Generator{}

type VideoController interface {
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	videos []models.Video
}

func New() VideoController {
	return &controller{}
}

func (c *controller) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.videos)
}

func (c *controller) Update(ctx *gin.Context) {
	var uVideo models.Video
	if err := ctx.ShouldBind(&uVideo); err != nil {
		ctx.String(http.StatusBadRequest, "bad request %v", err)
		return
	}
	if err := ctx.ShouldBindUri(&uVideo); err != nil {
		ctx.String(http.StatusBadRequest, "bad request %v", err)
		return
	}
	for i, video := range c.videos {
		if video.Id == uVideo.Id {
			c.videos[i] = uVideo
			ctx.String(http.StatusOK, "video id %d has been updated", uVideo.Id)
			return
		}
	}
	ctx.String(http.StatusBadRequest, "video id %d not found", uVideo.Id)
}

func (c *controller) Create(ctx *gin.Context) {
	var cVideo = models.Video{Id: generator.GetNextId()}
	if err := ctx.Bind(&cVideo); err != nil {
		ctx.String(http.StatusBadRequest, "bad request %v", err)
		return
	}
	c.videos = append(c.videos, cVideo)
	ctx.String(http.StatusOK, "create success new id is %d", cVideo.Id)
}

func (c *controller) Delete(ctx *gin.Context) {
	var dVideo models.Video
	if err := ctx.Bind(&dVideo); err != nil {
		ctx.String(http.StatusBadRequest, "bad request %v", err)
		return
	}
	for i, video := range c.videos {
		if video.Id == dVideo.Id {
			c.videos = append(c.videos[:i], c.videos[i+1:]...)
			ctx.String(http.StatusOK, "id %d delete success", dVideo.Id)
			return
		}
	}
	ctx.String(http.StatusBadRequest, "bad request id %d not found", dVideo.Id)
}
