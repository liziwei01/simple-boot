/*
 * @Author: liziwei01
 * @Date: 2022-03-04 22:06:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-11-01 22:38:55
 * @Description: file content
 */
package bootstrap

import (
	"github.com/gin-gonic/gin"
)

// InitHandler 用*gin.Engine作http handler
func InitHandler(app *AppServer) *gin.Engine {
	gin.SetMode(app.Config.RunMode)
	handler := gin.Default()
	handler.ContextWithFallback = true
	return handler
}
