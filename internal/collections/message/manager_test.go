package message

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/chat_manager"
	"github.com/mahdi-cpp/contacts-api/internal/collections/chat"
	"github.com/mahdi-cpp/contacts-api/internal/config"
	"github.com/mahdi-cpp/contacts-api/internal/help"
)

func TestCreateMessage(t *testing.T) {

	config.Init()

	ch := &chat.Chat{
		ID:    config.ChatID1,
		Title: "Sara",
		Type:  "private",
		MembersCount: []chat.Member{
			{
				UserID: config.Mahdi,
				Role:   "admin",
			},
			{
				UserID: config.Ali,
				Role:   "admin",
			},
			{
				UserID: config.Golnar,
				Role:   "member",
			},
		},
	}

	manager, err := chat_manager.New(ch)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		msg := &Message{
			ChatID:    ch.ID,
			Caption:   "sara " + strconv.Itoa(i),
			UserID:    uuid.New(),
			Type:      "sound",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = manager.CreateMessage(msg)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestReadMessage(t *testing.T) {

	config.Init()
	ch := &chat.Chat{
		ID:    config.ChatID1,
		Title: "chat with ali",
		Type:  "private",
		MembersCount: []chat.Member{
			{
				UserID:   config.Mahdi,
				Role:     "admin",
				IsActive: true,
				JoinedAt: time.Now(),
			},
			{
				UserID:   config.Ali,
				Role:     "admin",
				IsActive: true,
				JoinedAt: time.Now(),
			},
			{
				UserID:   config.Golnar,
				Role:     "member",
				IsActive: true,
				JoinedAt: time.Now(),
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	manager, err := chat_manager.New(ch)
	if err != nil {
		t.Fatal(err)
	}

	start := time.Now()
	with := &SearchOptions{
		//Type:      "photo",
		SortOrder: "end",
		Sort:      "createdAt",
		Page:      1,
		Size:      10,
	}

	msgs, err := manager.ReadMessages(with)
	if err != nil {
		t.Fatal(err)
	}

	for _, msg := range msgs {
		log.Println("message-->", msg.Caption)
	}
	dur := time.Since(start)
	fmt.Println(dur)
}

func TestUpdateMessage(t *testing.T) {

	config.Init()
	ch := &chat.Chat{
		ID:    config.ChatID1,
		Title: "chat with ali",
		Type:  "private",
		MembersCount: []chat.Member{
			{
				UserID: config.Mahdi,
				Role:   "admin",
			},
			{
				UserID: config.Ali,
				Role:   "admin",
			},
			{
				UserID: config.Golnar,
				Role:   "member",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	manager, err := chat_manager.New(ch)
	if err != nil {
		t.Fatal(err)
	}

	id, err := uuid.Parse("019968b8-1a9a-7267-861d-5ba3428c8044")
	if err != nil {
		t.Fatal(err)
	}

	with := &UpdateOptions{
		ID:       id,
		Caption:  help.StrPtr("Dad 09125640293"),
		IsPinned: help.BoolPtr(true),
	}

	_, err = manager.UpdateMessage(with)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteMessage(t *testing.T) {
	config.Init()
	ch := &chat.Chat{
		ID:    config.ChatID1,
		Title: "chat with ali",
		Type:  "private",
		MembersCount: []chat.Member{
			{
				UserID: config.Mahdi,
				Role:   "admin",
			},
			{
				UserID: config.Ali,
				Role:   "admin",
			},
			{
				UserID: config.Golnar,
				Role:   "member",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	manager, err := chat_manager.New(ch)
	if err != nil {
		t.Fatal(err)
	}

	id, err := uuid.Parse("019968b8-1a9a-72bb-aaa6-6aae3ded0957")
	if err != nil {
		t.Fatal(err)
	}

	err = manager.DeleteMessage(id)
	if err != nil {
		t.Fatal(err)
	}
}
