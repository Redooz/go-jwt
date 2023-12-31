package controllers

import "github.com/gin-gonic/gin"

// HelloUser godoc
// @Summary Greet the user
// @Description Returns a greeting to the user
// @Tags user
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string "Successfully greeted the user"
// @Failure 401
// @Router /user/hello [get]
func HelloUser(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello User",
    })
}
