package member

import (
	"fmt"
	"testing"

	"github.com/mahdi-cpp/contacts-api/internal/config"
)

const path = "/app/iris/com.iris.messages/metadata"

func Test_ReadMembers(t *testing.T) {

	config.Init()
	manager, err := NewManager(path, "members")
	if err != nil {
		t.Fatal(err)
	}

	with := &SearchOptions{
		//IsPin: help.BoolPtr(true),
	}

	items, err := manager.ReadWith(with)
	if err != nil {
		return
	}

	for _, item := range items {
		fmt.Println(item.IsPin)
	}
}

func Test_Clone(t *testing.T) {

	config.Init()
	manager, err := NewManager(path, "members")
	if err != nil {
		t.Fatal(err)
	}

	err = manager.Copy(path, "members_v2")
	if err != nil {
		t.Fatal(err)
	}
}
