package router

import (
	"VueGin/global"
	authhandler "VueGin/handler/auth"
	"VueGin/repository"
	"VueGin/service"

	"github.com/gin-gonic/gin"
)

func InitLoginRouter(Router *gin.RouterGroup) {
	methods := authhandler.AuthHandler{
		AuthSrv: &service.AuthService{
			Repo: &repository.AuthRepository{
				DB: global.Global_DB,
			},
		},
	}
	login := Router.Group("/login")
	{
		login.POST("/auth", methods.AuthByJWT)
	}
}
