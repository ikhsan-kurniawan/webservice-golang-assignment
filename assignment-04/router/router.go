package router

import (
	"mygram/controllers"
	"mygram/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartApp(db *gorm.DB) *gin.Engine {
	userRepository := repository.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)
	_ = userController

	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userController.UserRegister)
		userRouter.POST("/login", userController.UserLogin)
	}

	// productRouter := r.Group("/products")
	// {
	// 	productRouter.Use(middlewares.Authentication())
	// 	productRouter.POST("/", controllers.CreateProduct)

	// 	productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	// }

	return r
}