package group

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/collections/contact"
)

const (
	userID  = "018fe65d-8e4a-74b0-8001-c8a7c29367e1"
	groupID = "0199860a-7364-7de0-8ce5-f1a666df77a5"
	workDir = "/app/iris/com.iris.contacts/users/018fe65d-8e4a-74b0-8001-c8a7c29367e1/metadata"
	name    = "contacts"
)

func TestCreatePhotos(t *testing.T) {

	contactManager, err := contact.NewManager(workDir, name)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		p := &contact.Contact{
			FirstName: "Ali",
			LastName:  "Abdolmaleki",
			Phones: []contact.Phone{
				{
					Value: "09355512617",
					Label: "Mobile 1",
				},
				{
					Value: "02145698789",
					Label: "Home 1",
				},
			},
			OriginalURL:  "photo1.jpg",
			ThumbnailURL: "photo1_270.jpg",
		}

		_, err := contactManager.Create(p)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCreateGroup(t *testing.T) {
	contactManager, err := contact.NewManager(workDir, "contacts")
	if err != nil {
		t.Fatal(err)
	}

	groupManager, err := NewManager(contactManager, workDir)
	if err != nil {
		t.Fatal(err)
	}

	a := &Group{
		Title:    "Group 3",
		IsHidden: true,
		Version:  "1.0.0",
	}

	create, err := groupManager.Create(a)
	if err != nil {
		return
	}
	fmt.Println(create.ID)
}

func TestAddContact(t *testing.T) {

	contactManager, err := contact.NewManager(workDir, name)
	if err != nil {
		t.Fatal(err)
	}

	groupManager, err := NewManager(contactManager, workDir)
	if err != nil {
		t.Fatal(err)
	}

	groupID, err := uuid.Parse(groupID)
	if err != nil {
		return
	}
	group, err := groupManager.Read(groupID)
	if err != nil {
		return
	}

	in := contactManager.ReadIndexes()
	for _, i := range in {
		fmt.Println(i.ID)
		err := groupManager.AddContact(group.ID, i.ID)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGroup(t *testing.T) {

	contactManager, err := contact.NewManager(workDir, name)
	if err != nil {
		t.Fatal(err)
	}

	groupManager, err := NewManager(contactManager, workDir)
	if err != nil {
		t.Fatal(err)
	}

	//all, err := groupManager.ReadAll()
	//if err != nil {
	//	return
	//}
	//for _, i := range all {
	//	fmt.Println(i.ID)
	//}

	a, err := uuid.Parse(groupID)
	if err != nil {
		t.Fatal(err)
	}

	with := &contact.SearchOptions{
		Sort:      "id",
		SortOrder: "desc",
	}

	//all, err := groupManager.ReadGroups(with)
	//if err != nil {
	//	t.Fatal(err)
	//}

	//for _, p := range all.Photos {
	//	fmt.Println(p.FileInfo.OriginalURL)
	//}
}

//func TestGroupPhotosLimit(t *testing.T) {
//
//	contactManager, err := contact.NewManager(workDir, name)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	groupManager, err := NewManager(contactManager, workDir)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	start := time.Now()
//
//	with := &contact.SearchOptions{
//		Sort:      "id",
//		SortOrder: "desc",
//		Page:      1,
//		Size:      2,
//	}
//
//	groups := groupManager.ReadGroups(with)
//
//	for _, item := range groups {
//		fmt.Println(item.Contacts)
//	}
//
//	duration := time.Since(start)
//	fmt.Println(duration)
//}
