package router

import (
	"mygram/controllers"
	"mygram/middlewares"
	"mygram/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartApp(db *gorm.DB) *gin.Engine {
	userRepository := repository.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	photoRepository := repository.NewPhotoRepository(db)
	photoController := controllers.NewPhotoController(photoRepository)

	commentRepository := repository.NewCommentRepository(db)
	commentController := controllers.NewCommentController(commentRepository)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaController := controllers.NewSocialMediaController(socialMediaRepository)



	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userController.UserRegister)
		userRouter.POST("/login", userController.UserLogin)

		userRouter.PUT("/:userId", middlewares.Authentication(), middlewares.UserAuthorization(), userController.UpdateUser)
		userRouter.DELETE("", middlewares.Authentication(), userController.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("", photoController.CreatePhoto)
		photoRouter.GET("", photoController.GetPhotos)
		photoRouter.GET("/:photoId", middlewares.PhotoAuthorization(db), photoController.GetPhotoById)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(db), photoController.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(db), photoController.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("", commentController.CreateComment)
		commentRouter.GET("", commentController.GetComments)
		commentRouter.GET("/:commentId", middlewares.CommentAuthorization(db), commentController.GetCommentById)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(db), commentController.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(db), commentController.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("", socialMediaController.CreateSocialMedia)
		socialMediaRouter.GET("", socialMediaController.GetSocialMedias)
		socialMediaRouter.GET("/:socialMediaId", middlewares.SocialMediaAuthorization(db), socialMediaController.GetSocialMediaById)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(db), socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(db), socialMediaController.DeleteSocialMedia)
	}
	// productRouter := r.Group("/products")
	// {
	// 	productRouter.Use(middlewares.Authentication())
	// 	productRouter.POST("/", controllers.CreateProduct)

	// 	productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	// }

	return r
}