package controllers

import (
	"gin_blogs/err_response"
	"gin_blogs/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignupPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"users/signup.tpl",
		gin.H{})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"users/login.tpl",
		gin.H{})
}

type formData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Signup(c *gin.Context) {
	var data formData
	c.Bind(&data)

	// Check if the user exists already
	if !models.CheckUserAvailability(data.Email) {
		err_response.SendErrorResponse(c, http.StatusBadRequest,
			err_response.HTML,
			err_response.Response{Message: "Email missing"})
		return
	}

	// Create the user
	user := models.UserCreate(data.Email, data.Password)
	if user == nil || user.ID == 0 {
		err_response.SendErrorResponse(c, http.StatusBadRequest,
			err_response.HTML,
			err_response.Response{Error: "user creation failed"})
		return
	}
	// Set the session.
	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Save()
	c.Redirect(http.StatusFound, "/blogs")
}

func Login(c *gin.Context) {
	var data formData
	c.Bind(&data)

	// Match password
	user := models.UserMatchPassword(data.Email, data.Password)
	if user.ID == 0 {
		err_response.SendErrorResponse(c, http.StatusUnauthorized,
			err_response.HTML,
			err_response.Response{Error: "Unauthorized!"})
		return
	}
	// Set the session.
	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/blogs")
}

func Logout(c *gin.Context) {
	// Delete the session
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Status(http.StatusAccepted)
}
