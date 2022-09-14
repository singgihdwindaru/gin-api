package routers

import (
	controller "gin-api/src/controllers"
	"gin-api/src/models"
	"gin-api/src/repository/mysql"
	"gin-api/src/usecase"

	"github.com/gin-gonic/gin"
)

var (
	usersMysqlRepo models.IUserRepository

	userUsecase models.IUserUsecase
)

func initRepo() {
	usersMysqlRepo = mysql.NewUserMysqlRepository(nil)
}

func initUsecase() {
	userUsecase = usecase.NewUserUsecase(usersMysqlRepo)
}
func SetupRouter() *gin.Engine {
	initRepo()
	initUsecase()
	r := gin.Default()
	controller.NewUserController(r, userUsecase)
	return r
}
