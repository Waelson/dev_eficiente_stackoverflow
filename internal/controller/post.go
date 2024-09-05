package controller

import (
	"github.com/Waelson/internal/model"
	"github.com/Waelson/internal/service"
	"github.com/Waelson/internal/util/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController interface {
	Save(c *gin.Context)
}

type postController struct {
	postService service.PostService
}

func (a *postController) Save(c *gin.Context) {
	post := model.Post{}

	if !middleware.ValidateStruct(c, &post) {
		return
	}

	user := c.Request.Header.Get("user")
	post.User = user

	_, err := a.postService.Save(c.Request.Context(), &post)
	if err != nil {
		c.AbortWithError(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func NewPostController(postService service.PostService) PostController {
	result := postController{
		postService: postService,
	}
	return &result
}
