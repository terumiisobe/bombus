package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/terumiisobe/bombus/api/controllers"
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
