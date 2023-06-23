package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	routerEngine := gin.Default()
	routerEngine.Use(AuthMiddleware())
	/*
		grouping: auth
		www.enigmacamp.com/auth/login
		www.enigmacamp.com/auth/register
		www.enigacamp.com/auth/logout

		grouping: master
		www.enigmacamp.com/master/users
		www.enigmacamp.com/master/students

		www.enigmacamp.com/api/auth/login
	*/
	routerEngine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy Check")
	})
	rgAuth := routerEngine.Group("/api/auth")
	rgAuth.POST("/login", login)

	rgMaster := routerEngine.Group("/api/master")
	rgMaster.GET("/greeting/:name", greeting)

	err := routerEngine.Run() //bisa costum Run("...")
	if err != nil {
		panic(err)
	}
	//secara default menggunakan port: 8080
}

//func greeting(c *gin.Context) {
// 	name := c.Param("name") //withparameter :name at path/greeting
// 	c.String(200, "Hello...%s", name)
// }

func greeting(c *gin.Context) {
	name := c.Param("name")
	kec := c.Query("kec")
	kel := c.Query("kel")
	c.String(http.StatusOK, "Hello %s saat ini kamu berada di Kecamatan %s Kelurahan %s", name, kec, kel)
}

//func login(c *gin.Context) {
// 	username := c.PostForm("usename")
// 	c.PostForm("password")
// 	c.String(http.StatusOK "Hello %s", username)
// }

/*
Model Binding & Validation
-> ShouldBind, ShouldBindJSON -> using struct tag binding:required
*/

func login(c *gin.Context) {
	var uc UserCrendetial
	if err := c.BindJSON(&uc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token":    "ini_token",
		})
	}
}

// func login(c *gin.Context) {
// 	var uc UserCrendetial
// 	if c.ShouldBind(&uc) = nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "OK",
// 			"data": uc.Usernme,
// 		})
// 	}
// }