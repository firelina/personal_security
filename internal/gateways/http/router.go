package http

import (
	"github.com/gin-gonic/gin"
	"personal_security/internal/gateways/http/handlers"
)

func setupRouter(r *gin.Engine, useCases UseCases) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	userHandler := handlers.NewUserHandler(useCases.User)
	eventHandler := handlers.NewEventHandler(useCases.Event)
	contactHandler := handlers.NewContactHandler(useCases.Contact)
	reminderHandlers := handlers.NewReminderHandler(useCases.Reminder)

	r.POST("/users", func(c *gin.Context) {
		userHandler.RegisterUser(c)
	})

	r.POST("/login", func(c *gin.Context) {
		userHandler.LoginUser(c)
	})

	r.POST("/events", func(c *gin.Context) {
		eventHandler.CreateEvent(c)
	})

	r.POST("/contacts", func(c *gin.Context) {
		contactHandler.CreateEvent(c)
	})

	r.POST("/reminders", func(c *gin.Context) {
		reminderHandlers.CreateReminder(c)
	})

	r.GET("/events/:user_id", func(c *gin.Context) {
		eventHandler.GetEvents(c)
	})

	r.GET("/contacts/:user_id", func(c *gin.Context) {
		contactHandler.GetContacts(c)
	})

	r.GET("/reminders/:event_id", func(c *gin.Context) {
		reminderHandlers.GetReminders(c)
	})

	r.POST("/events/:id/update", func(c *gin.Context) {
		eventHandler.UpdateEventStatus(c)
	})

	r.POST("/reminders/remind/:user_id", func(c *gin.Context) {
		reminderHandlers.SendReminders(c)
	})
}
