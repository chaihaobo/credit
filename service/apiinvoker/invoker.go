// Package apiinvoker
// @author： Boice
// @createTime：2022/11/29 18:53
package apiinvoker

import "github.com/gin-gonic/gin"

type Invoker interface {
	Invoke(c *gin.Context) (any, error)
}
