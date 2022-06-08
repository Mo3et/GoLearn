// Link :https://github.com/lydell/resolve-url#deprecated
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	main2()
}

// Using GET POST PUT PATCH DELETE OPTIONS
// func main1() {
// 	// Creates a gin router with default middleware:
// 	// logger and recovery (crash-free) middleware
// 	router := gin.Default()

// 	router.GET("/someGet", getting)
// 	router.POST("/somePost", posting)
// 	router.PUT("/someput", putting)
// 	router.DELETE("/someDelete", deleting)
// 	router.PATCH("/somePatch", patching)
// 	router.HEAD("/someHead", head)
// 	router.OPTIONS("/someOptions", options)

// 	// By default it serves on :8000 unless a
// 	//PORT enviroment  variable was defined.
// 	router.Run(":3001")
// 	// router.Run()
// 	// router.Run(":3000") for a hard coded port
// }

// Parameters in path
func main2() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

	// This handler will add a new router for /user/groups.
	// Exact routes are resolved before param routes, regardless of the order they were defined.
	// Routes starting with /user/groups are never interpreted as /user/:name/... routes
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

	router.Run(":8080")
}

// Querystring parameters
func main3() {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching : /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") //shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8081")
}

// Multipart / Urlencoded Form
func main4() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8081")
}

// Map as querystring or postform parameters

// POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
//
// names[first]=thinkerou&names[second]=tianou

func main5() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v;names: %v", ids, names)
	})
	router.Run(":8081")
} // ids: map[b:hello a:1234]; names: map[second:tianou first:thinkerou]
