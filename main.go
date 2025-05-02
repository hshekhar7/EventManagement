package main

import (
	"fmt"

	"github.com/ank809/Event-Management-golang~/authentication"
	"github.com/ank809/Event-Management-golang~/controllers"
	"github.com/ank809/Event-Management-golang~/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "HELLO WORLD")
	})
	r.POST("/signup", authentication.SignUp)
	r.GET("/loginattendee", authentication.LoginAttendee)
	r.GET("/loginmanager", authentication.EventManagerLogin)
	authenticatedroutes := r.Group("/")
	authenticatedroutes.Use(middlewares.AuthMiddleware())
	{
		authenticatedroutes.GET("/readevents", controllers.ReadAllEvents)
		authenticatedroutes.GET("/registerevent/:id", controllers.RegisterEvent)
	}
	protectedeventroutes := r.Group("/events")
	protectedeventroutes.Use(middlewares.AuthMiddleware())
	protectedeventroutes.Use(middlewares.Authorize("EventManager"))
	{
		protectedeventroutes.POST("/addevent", controllers.AddEvent)
		protectedeventroutes.GET("/deleteevent/:id", controllers.DeleteEvents)
		protectedeventroutes.POST("/updateevent/:id", controllers.UpdateEvent)
	}
	if err := r.Run(":8081"); err != nil {

		fmt.Println(err)
	}

}
