package controller

import (
	"github.com/Waelson/internal/model"
	"github.com/Waelson/internal/service"
	"github.com/Waelson/internal/util/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnswerController interface {
	Save(c *gin.Context)
}

type answerController struct {
	answerService service.AnswerService
}

func (a *answerController) Save(c *gin.Context) {
	answer := model.Answer{}

	if !middleware.ValidateStruct(c, &answer) {
		return
	}

	user := c.Request.Header.Get("user")
	answer.User = user

	_, err := a.answerService.Save(c.Request.Context(), &answer)
	if err != nil {
		c.AbortWithError(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func NewAnswerController(answerService service.AnswerService) AnswerController {
	result := answerController{
		answerService: answerService,
	}
	return &result
}
