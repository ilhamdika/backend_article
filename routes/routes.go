package routes

import (
	"backend_article/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/article", controllers.CreateArticle)
	r.GET("/articles", controllers.GetArticles) 
	r.GET("/articles/:limit/:offset", controllers.GetArticles)
	r.GET("/article/:id", controllers.GetArticle)
	r.PUT("/article/:id", controllers.UpdateArticle)
	r.DELETE("/article/:id", controllers.DeleteArticle)
	r.PUT("/article-status/:id", controllers.UpdateStatus)
}
