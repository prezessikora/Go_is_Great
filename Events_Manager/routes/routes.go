package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById) //events/1 events/2
	server.POST("/events", createEvent)
	// delete
	// update

}
