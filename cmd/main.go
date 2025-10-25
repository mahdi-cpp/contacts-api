package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mahdi-cpp/contacts-api/internal/api/handlers"
	"github.com/mahdi-cpp/contacts-api/internal/application"
	"github.com/mahdi-cpp/contacts-api/internal/config"
	"github.com/mahdi-cpp/iris-tools/mygin"
)

func main() {

	config.Init()
	appManager, _ := application.New()

	chatHandler := handlers.NewChatHandler(appManager)
	messageHandler := handlers.NewMessageHandler(appManager)

	// Create a new engine with default middleware
	router := mygin.New()

	router.POST("/api/chats", chatHandler.Create)
	router.POST("/api/chats/search", chatHandler.ReadAll)

	router.PATCH("/api/members/", chatHandler.UpdateMember)

	//router.GET("/api/chats", chatHandler.Read)

	// Register routes
	router.POST("/api/messages", messageHandler.Create)
	router.GET("/api/messages", messageHandler.Read2)
	//router.GET("/api/messages/:id", messageHandler.Read)
	//router.PATCH("/api/messages/:id", messageHandler.Update)
	//router.DELETE("/api/messages/:id", messageHandler.Delete)

	// Start the server
	fmt.Println("Server is running on http://localhost:50152")
	log.Fatal(http.ListenAndServe(":50152", router))
}
