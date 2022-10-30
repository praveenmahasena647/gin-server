package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/vue-todo/dbs"
)

func RunServer() {
	var router = gin.Default()

	router.GET("/", getAll)
	router.GET("/:id", getOne)
	router.POST("/", postOne)
	router.PUT("/:id", editOne)
	router.DELETE("/:id", deleteOne)
	router.DELETE("/", deleteAll)

	router.Run("localhost:42069")

}

func getAll(c *gin.Context) {
	var todo []dbs.ToDo
	dbs.Db.Find(&todo)
	c.JSONP(200, todo)
}
func getOne(c *gin.Context) {
	var id = c.Param("id")
	log.Println(id)
}
func postOne(c *gin.Context) {
	c.JSONP(200, map[string]string{
		"msg": "post One",
	})
}
func editOne(c *gin.Context) {
	c.JSONP(200, map[string]string{
		"msg": "editOne",
	})
}
func deleteOne(c *gin.Context) {
	c.JSONP(200, map[string]string{
		"msg": "deleteOne",
	})
}
func deleteAll(c *gin.Context) {
	c.JSONP(200, map[string]string{
		"msg": "delete all",
	})
}
