package controllers

import (
	"gin_blogs/err_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	id := c.Query("id")

	if len(id) == 0 {
		err_response.SendErrorResponse(c,
			http.StatusBadRequest,
			err_response.HTML,
			err_response.Response{Message: "id param missing!"},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All Okay!"})
}
