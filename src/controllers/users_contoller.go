package controller

import (
	"gin-api/src/models"
	"gin-api/src/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userCtrl struct {
	userUsecase models.IUserUsecase
}

func NewUserController(r *gin.Engine, userusecase models.IUserUsecase) {
	c := &userCtrl{
		userUsecase: userusecase,
	}
	r.POST("/signin", c.SignIn)
}

func (c *userCtrl) CreateUser(g *gin.Context) {
	request := models.UserRequest{}

	if err := g.BindJSON(&request); err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, utils.HttpResponse(http.StatusBadRequest, "Invalid request", nil))
		return
	}

	result, err := c.userUsecase.CreateUser(g.Request.Context(), request)
	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusInternalServerError, utils.HttpResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	g.JSON(http.StatusCreated, utils.HttpResponse(http.StatusAccepted, "Success create user", result))

}
func (c *userCtrl) SignIn(g *gin.Context) {
	request := models.LoginRequest{}

	if err := g.BindJSON(&request); err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, utils.HttpResponse(http.StatusBadRequest, "Invalid request", nil))
		return
	}

	result, err := c.userUsecase.Login(g.Request.Context(), request)
	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusInternalServerError, utils.HttpResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	g.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, "Success SignIn", result))
	// return
}
