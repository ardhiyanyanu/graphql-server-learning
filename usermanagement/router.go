package usermanagement

import (
	"github.com/gin-gonic/gin"
	"github.com/alterra/graphql-server/usermanagement/application"
)

func GetRoute(r *gin.RouterGroup) {
	// r.POST("/customer/register", application.RegisterNewCustomerAndUser)
	r.POST("/login", application.Login)
}
