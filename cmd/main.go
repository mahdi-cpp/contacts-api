package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mahdi-cpp/contacts-api/internal/api/contacts"
	group_handler "github.com/mahdi-cpp/contacts-api/internal/api/groups"
	"github.com/mahdi-cpp/contacts-api/internal/application"
	"github.com/mahdi-cpp/contacts-api/internal/collections/contact"
	"github.com/mahdi-cpp/contacts-api/internal/config"
	"github.com/mahdi-cpp/iris-tools/mygin"
)

func main() {

	config.Init()
	appManager, _ := application.New()

	manager, err := appManager.GetAccountManager(config.Mahdi)
	if err != nil {
		return
	}

	with := &contact.SearchOptions{}

	all, err := manager.ContactManager.ReadAll(with)
	if err != nil {
		return
	}
	fmt.Println("contacts count: ", len(all))

	// Create a new engine with default middleware
	router := mygin.New()

	photoHandler := contact_handler.New(appManager)
	albumHandler := group_handler.New(appManager)

	router.POST("/api/contacts", photoHandler.Create)
	router.POST("/api/contacts/search", photoHandler.ReadAll)

	router.GET("/api/contacts", photoHandler.Read)

	router.PATCH("/api/contacts", photoHandler.Update)

	//router.DELETE("/api/contacts", photoHandler.Delete)

	//---------------------------------------------------------------------------
	router.POST("/api/groups", albumHandler.Create)
	router.POST("/api/groups/contacts", albumHandler.AddContact)
	router.POST("/api/groups/search", albumHandler.ReadGroups)
	router.POST("/api/groups/ali", albumHandler.ReadGroupContacts)
	router.DELETE("/api/groups", albumHandler.Delete)

	// Start the server
	fmt.Println("Server is running on http://localhost:50153")
	log.Fatal(http.ListenAndServe(":50153", router))
}
