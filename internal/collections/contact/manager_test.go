package contact

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/config"
)

func handlePhotoCreation(message string, id uuid.UUID) {
	switch message {
	case "create":

	case "update":

	case "delete":

	}
}

func Test_Create(t *testing.T) {

	config.Init()
	var assets = "/app/iris/com.iris.contacts/users/" + config.Mahdi.String() + "/thumbnails/"

	manager, err := NewManager(config.Mahdi, handlePhotoCreation, "/app/iris/com.iris.contacts/users/"+config.Mahdi.String()+"/metadata")
	if err != nil {
		t.Fatal(err)
	}

	//co := &Contact{
	//	FirstName:    "Negin Noor Mohammadi",
	//	LastName:     "Noor Mohammadi",
	//	OriginalURL:  "",
	//	ThumbnailURL: assets + "chat_35",
	//}

	co := &Contact{
		FirstName:    "Dariush Forouhar",
		LastName:     "Noor Mohammadi",
		OriginalURL:  "",
		ThumbnailURL: assets + "chat_17",
	}

	contact, err := manager.Create(co)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(contact.ID)
}
