package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevindharmawan/saas-backend/internal/model"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func SendSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Data:  data,
		Error: nil,
	})
}

func SendErrorResponse(c *gin.Context, err *model.AppError) {
	c.AbortWithStatusJSON(
		err.Status(),
		Response{
			Data:  nil,
			Error: err.Error(),
		},
	)
}
