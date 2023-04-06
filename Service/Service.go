package Service

import (
	"github.com/gin-gonic/gin"
	"github.com/itcyy/Database"
	"net/http"
)

/*
*   Run
*   <运行整个项目的运行函数>
*   <API接口>
*   @author: [陈永裕]
*   @version[v1.0.0.1,2023-4-1]
*   @param []
*   @Description:
 */
func Run() {
	r := gin.Default()

	r.Use(Cors())
	r.POST("/test", func(c *gin.Context) {
		data := Database.SelectData("select * from data")
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	r.GET("/sreach", func(c *gin.Context) {
		name := c.Query("name")

		date := Database.SelectData("select * from data where bookname like '" + name + "%'")
		c.JSON(http.StatusOK, gin.H{
			"data": date,
		})
	})
	r.RunTLS(":9090", "/etc/pki/nginx/itcyy.cn_bundle.crt", "/etc/pki/nginx/itcyy.cn.key")
}

/*
*   Cors
*   @author: [陈永裕]
*   @version[v1.0.0.1,2023-4-1]
*   @Description:
*   @return [gin.HandlerFunc]
 */
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
