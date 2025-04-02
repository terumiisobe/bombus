package routes

import (
	"github.com/terumiisobe/bombus/api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	colmeiaRoutes := r.Group("/colmeia")
	{
		colmeiaRoutes.GET("/", controllers.GetColmeias)
		colmeiaRoutes.GET("/:id", controllers.GetColmeia)
		colmeiaRoutes.POST("/", controllers.CreateColmeia)
		//colmeiaRoutes.UPDATE("/", controllers.UpdateColmeia)
		colmeiaRoutes.DELETE("/:id", controllers.DeleteColmeia)
	}
}
