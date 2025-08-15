// routes/routes.go
package routes

import (
	"ternago/backend-api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine) {
	// routes for health check
	r.GET("/health", controllers.HealthCheck)
	api := r.Group("/api/v1")
	{
		// AuthMstAdmins routes
		admins := api.Group("/admins")
		{
			admins.POST("", controllers.CreateAuthMstAdmin)
			admins.GET("/all", controllers.GetAuthMstAdmins)
			admins.GET("/:id", controllers.GetAuthMstAdmin)
			admins.PUT("/:id", controllers.UpdateAuthMstAdmin)
			admins.DELETE("/:id", controllers.DeleteAuthMstAdmin)
		}
	}
}
