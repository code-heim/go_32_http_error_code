package err_response

import (
	"github.com/gin-gonic/gin"
)

type ResponseType string

const (
	HTML ResponseType = "HTML"
	JSON ResponseType = "JSON"
)

type Response struct {
	Message string
	Error   string
}

func SendErrorResponse(c *gin.Context, status int, respType ResponseType, response Response) {
	data := gin.H{
		"Message": response.Message,
		"Error":   response.Error,
	}

	switch respType {
	case HTML:
		c.HTML(status, "errors/error.tpl", data)
	case JSON:
		c.JSON(status, data)
	default:
		c.HTML(status, "errors/error.tpl", data)
	}
}
