package chat

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/collections/member"
	"github.com/mahdi-cpp/contacts-api/internal/config"
	"github.com/mahdi-cpp/contacts-api/internal/help"
)

const chatDatabaseName = "chats"

func TestCreateChat(t *testing.T) {

	config.Init()
	manager, err := NewManager("/app/iris/com.iris.messages/metadata", chatDatabaseName)
	if err != nil {
		t.Fatal(err)
	}

	ch := &Chat{
		Title:       "Fharhad",
		Type:        "private",
		OriginalURL: "",
		//ThumbnailURL: "/app/iris/services/accounts/assets/thumbnails/cat",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newChat, err := manager.CreateChat(ch)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("new chat id: ", newChat.ID)

	m1 := &member.Member{
		ChatID:      newChat.ID,
		UserID:      config.Nader,
		Role:        "member",
		CustomTitle: "CustomTitle",
		IsActive:    true,
		DeletedAt:   time.Now(),
	}

	m2 := &member.Member{
		ChatID:      newChat.ID,
		UserID:      config.Mahdi,
		Role:        "member",
		CustomTitle: "CustomTitle",
		IsActive:    true,
		DeletedAt:   time.Now(),
	}

	addMember1, err := manager.AddMember(m1)
	if err != nil {
		t.Fatal(err)
	}

	addMember2, err := manager.AddMember(m2)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("add member id: ", addMember1.ID)
	fmt.Println("add member id: ", addMember2.ID)
}

func TestManager_AddMember(t *testing.T) {
	config.Init()
	manager, err := NewManager("/app/iris/com.iris.messages/metadata", chatDatabaseName)
	if err != nil {
		t.Fatal(err)
	}

	chatId, err := uuid.Parse("0199ee21-e5b3-784a-8398-1bba4cab57d6")
	if err != nil {
		t.Fatal(err)
	}

	m2 := &member.Member{
		ChatID:      chatId,
		UserID:      config.Nader,
		Role:        "member",
		CustomTitle: "Maryam Add...",
		IsActive:    true,
		DeletedAt:   time.Now(),
	}

	newMember, err := manager.AddMember(m2)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("member id is: ", newMember.ID)
}

func TestReadChat(t *testing.T) {

	config.Init()
	manager, err := NewManager("/app/iris/com.iris.messages/metadata", chatDatabaseName)
	if err != nil {
		t.Fatal(err)
	}

	//with := &SearchOptions{
	//	UserID:    config.Mahdi,
	//	Sort:      "createdAt",
	//	SortOrder: "desc",
	//	Page:      1,
	//	Size:      100,
	//}

	fmt.Println("Ali: ")

	chats, err := manager.ReadUserChats(config.Mahdi)
	if err != nil {
		t.Fatal(err)
	}

	for _, ch := range chats {
		fmt.Println("---------: "+ch.Chat.Title, ch.Chat.ID)
	}
}

func Test_Update(t *testing.T) {

	config.Init()
	manager, err := NewManager("/app/iris/com.iris.messages/metadata", chatDatabaseName)
	if err != nil {
		t.Fatal(err)
	}

	chatId, err := uuid.Parse("0199ee21-e5b3-784a-8398-1bba4cab57d6")
	if err != nil {
		t.Fatal(err)
	}

	with := &UpdateOptions{
		ID:    chatId,
		Title: help.StrPtr("تیم بازاریابی و مارکتینگ"),
		//ThumbnailURL: help.StrPtr("/app/iris/services/accounts/assets/thumbnails/irancell3"),
	}

	err = manager.UpdateChat(with)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteChat(t *testing.T) {
	config.Init()
	manager, err := NewManager("/app/iris/com.iris.messages/metadata", chatDatabaseName)
	if err != nil {
		t.Fatal(err)
	}

	chatId, err := uuid.Parse("0199fd93-6ad7-732b-8e3e-9b9ca7b0d545")
	if err != nil {
		t.Fatal(err)
	}

	err = manager.ChatDelete(chatId)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCreateChats(t *testing.T) {

	manager, err := NewManager("/app/iris/com.iris.messages/metadata", chatDatabaseName)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {

		userID1, err := uuid.NewV7()
		if err != nil {
			t.Fatal(err)
		}
		userID2, err := uuid.NewV7()
		if err != nil {
			t.Fatal(err)
		}

		ch := &Chat{
			Username:     fmt.Sprintf("status%d", i),
			Type:         "bot",
			Version:      "2",
			Description:  "mahdi bot" + strconv.Itoa(i),
			ThumbnailURL: "avatar_" + strconv.Itoa(i),
			CreatedAt:    time.Now(),
		}

		_, err = manager.CreateChat(ch)
		if err != nil {
			t.Fatal(err)
		}

		Members := []*member.Member{
			{
				UserID:    userID1,
				Role:      "member",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    userID2,
				Role:      "member",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    userID1,
				Role:      "member",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    userID2,
				Role:      "member",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    userID1,
				Role:      "member",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    userID2,
				Role:      "member",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    userID1,
				Role:      "member",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    userID2,
				Role:      "member",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    userID1,
				Role:      "member",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    userID2,
				Role:      "member",
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		for _, mem := range Members {
			_, err := manager.MemberManager.Create(mem)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

//
//func TestReadChats(t *testing.T) {
//
//	db, err := collection_manager_memory.New[*Chat](databaseDirectory)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer db.Close()
//
//	all, err := db.ReadAll()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	sh := &SearchOptions{
//		Type:      help.StrPtr("private"),
//		SortOrder: "end",
//		Sort:      "createdAt",
//	}
//
//	filterItems := Search(all, sh)
//
//	for _, ch := range filterItems {
//		fmt.Println(ch.Description)
//	}
//}
//
//func TestUpdate(t *testing.T) {
//	db, err := collection_manager_memory.New[*Chat](databaseDirectory)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer db.Close()
//
//	all, err := db.ReadAll()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	updateOptions := &UpdateOptions{
//		Type: help.StrPtr("private"),
//	}
//
//	for _, ch := range all {
//		Update(ch, updateOptions)
//		_, err := db.Update(ch)
//		if err != nil {
//			t.Fatal(err)
//		}
//	}
//}

const path = "/app/iris/com.iris.messages/metadata"

func Test_Copy(t *testing.T) {

	config.Init()
	manager, err := NewManager(path, "chats")
	if err != nil {
		t.Fatal(err)
	}

	err = manager.Clone(path, "chats_v2")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ReadChats(t *testing.T) {

	config.Init()
	manager, err := NewManager("/app/iris/com.iris.messages/metadata", "chats_v2")
	if err != nil {
		t.Fatal(err)
	}

	chats, err := manager.ReadUserChats(config.Mahdi)
	if err != nil {
		return
	}

	fmt.Printf("chats length: %v\n", len(chats))

	for _, chat := range chats {
		fmt.Println(chat.Chat.ID, chat.Chat.Title, chat.Chat.Type)
	}
}
