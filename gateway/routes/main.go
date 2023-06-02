package routes

import (
	"github.com/gin-gonic/gin"
)

// SetRoutes creates the http routes.
// If you want add a new one just create a new Group with the URI
func SetRoutes(r *gin.Engine) {
	sarg := r.Group("/summary-account")
	addSummaryAccountRoutes(sarg)
}
