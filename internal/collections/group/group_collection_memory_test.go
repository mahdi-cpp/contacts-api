package group

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/collections/contact"
	"github.com/mahdi-cpp/contacts-api/internal/config"
	"github.com/mahdi-cpp/iris-tools/collection_manager_join"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
)

func TestGroup_Create(t *testing.T) {

	path := config.GetUserMetadataPath("01997cba-6dab-7636-a1f8-2c03174c7b6e", "")
	alumCollection, _ := collection_manager_memory.New[*Group](path+"/groups", "groups")
	photoGroupsCollection, _ := collection_manager_join.New[*contact.Join](path+"/groups", "groups_join")

	a := &Group{
		Title:    "collection 1",
		Type:     "collection",
		IsHidden: false,
	}
	_, err := alumCollection.Create(a)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		contactID, err := uuid.NewV7()
		if err != nil {
			t.Fatal(err)
		}
		photoGroup := &contact.Join{
			ParentID:  a.ID,
			ContactID: contactID,
		}
		_, err = photoGroupsCollection.Create(photoGroup)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGroup_AddPhoto(t *testing.T) {

	path := config.GetUserMetadataPath("01997cba-6dab-7636-a1f8-2c03174c7b6e", "")

	photoGroupsCollection, _ := collection_manager_join.New[*contact.Join](path+"/groups", "groups_join")

	groupID, err := uuid.Parse("01997cba-6dab-7636-a1f8-2c03174c7b6e")
	if err != nil {
		t.Fatal(err)
	}

	p := &contact.Join{
		ParentID:  groupID,
		ContactID: uuid.New(),
	}

	_, err = photoGroupsCollection.Create(p)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGroup_ReadGroupPhotos(t *testing.T) {
	path := config.GetUserMetadataPath("01997cba-6dab-7636-a1f8-2c03174c7b6e", "")
	photoGroupsCollection, _ := collection_manager_join.New[*contact.Join](path+"/groups", "groups_join")

	groupID, err := uuid.Parse("01997cba-6dab-7636-a1f8-2c03174c7b6e")
	if err != nil {
		t.Fatal(err)
	}

	ids, err := photoGroupsCollection.GetByParentID(groupID)
	if err != nil {
		return
	}
	for _, id := range ids {
		fmt.Println("-------:", id.ContactID)
	}
}

func TestGroup_RemovePhoto(t *testing.T) {

}
