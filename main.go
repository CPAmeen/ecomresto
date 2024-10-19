// package main

// import (
// 	"ecomresto/controllers"
// 	"ecomresto/utils"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// var R = gin.Default()

// func init() {
// 	utils.InitDB()
// 	R.LoadHTMLGlob("views/*")

// }

// func main() {
// 	R.GET("/admin/login", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "admin_login.html", nil)
// 	})

// 	R.POST("/admin/login", controllers.AdminLogin)
// 	R.GET("/admin/home", func(c *gin.Context) {
// 		c.HTML(http.StatusSeeOther, "admin_home.html", nil)
// 	})
// 	R.GET("user/signup", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "signup.html", nil)
// 	})

// 	R.POST("user/signup", controllers.Signup)
// 	R.Run(":8080") // listen and serve on 0.0.0.0:8080

// }
package main

import (
	"ecomresto/controllers" // Import your controllers package
	"ecomresto/initializers"
	"ecomresto/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	initializers.ConnecttoDb()

	// Sync database (migrate schema)
	initializers.SyncDatabase(initializers.DB)
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil) // Make sure "signup.html" exists in your views directory
	})
	r.POST("/signup", controllers.Signup)
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil) // Make sure "login.html" exists in your views directory
	})
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	// After syncing the database
	r.Run(":5500")

}
