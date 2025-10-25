package handlers

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/mahdi-cpp/contacts-api/internal/application"
	"github.com/mahdi-cpp/contacts-api/internal/collections/chat"
	"github.com/mahdi-cpp/contacts-api/internal/collections/member"
	"github.com/mahdi-cpp/contacts-api/internal/help"
	"github.com/mahdi-cpp/iris-tools/mygin"
)

type ChatHandler struct {
	appManager *application.AppManager
}

func NewChatHandler(appManager *application.AppManager) *ChatHandler {
	return &ChatHandler{
		appManager: appManager,
	}
}

func (h *ChatHandler) Create(c *mygin.Context) {

	fmt.Println("ChatHandler create")

	var request *chat.Chat
	err := json.NewDecoder(c.Req.Body).Decode(&request)
	if err != nil {
		http.Error(c.Writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	newMessage, err := h.appManager.GetChatManager()
	if err != nil {
		// اینجا می‌توانید خطای مناسب را به کاربر برگردانید
		http.Error(c.Writer, "Failed to create chat", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(c.Writer).Encode(newMessage)
	if err != nil {
		return
	}
}

func (h *ChatHandler) ReadAll(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		fmt.Println("status id invalid")
		return
	}

	chatManager, err := h.appManager.GetChatManager()
	if err != nil {
		// اینجا می‌توانید خطای مناسب را به کاربر برگردانید
		http.Error(c.Writer, "Failed to create chat", http.StatusInternalServerError)
		return
	}

	chats, err := chatManager.ReadUserChats(userID)
	if err != nil {
		fmt.Println("2")
		http.Error(c.Writer, "Failed to read chats", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(c.Writer).Encode(chats)
	if err != nil {
		fmt.Println("3")
		return
	}

	//fmt.Println("chats: ", chats[2].Chat)
}

func (h *ChatHandler) UpdateMember(c *mygin.Context) {

	fmt.Println("ChatHandler updateMember")

	var with *member.UpdateOptions
	err := json.NewDecoder(c.Req.Body).Decode(&with)
	if err != nil {
		http.Error(c.Writer, "Invalid with body", http.StatusBadRequest)
		return
	}

	fmt.Println("with.IsPin: ", *with.IsPin)

	chatManager, err := h.appManager.GetChatManager()
	if err != nil {
		// اینجا می‌توانید خطای مناسب را به کاربر برگردانید
		http.Error(c.Writer, "Failed to create chat", http.StatusInternalServerError)
		return
	}

	err = chatManager.MemberManager.Update(with)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(c.Writer, "Failed to update chat", http.StatusInternalServerError)
		return
	}

	fmt.Println("8")

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusCreated)
}

//func (h *ChatHandler) Read(c *mygin.Context) {
//
//	//parse, err := uuid.Parse("01996b82-bcd8-7ca0-b148-e33aaac1fd85")
//	//if err != nil {
//	//	return
//	//}
//	fmt.Println("1")
//
//	with := &chat.SearchOptions{
//		Sort:      "createdAt",
//		SortOrder: "desc",
//	}
//	chats, err := h.appManager.ReadChats(with)
//	if err != nil {
//		fmt.Println("2")
//		http.Error(c.Writer, "Failed to read chats", http.StatusInternalServerError)
//		return
//	}
//
//	for _, ch := range chats {
//		fmt.Println(ch.ID, ch.Title)
//	}
//
//	c.Writer.Header().Set("Content-Type", "application/json")
//	c.Writer.WriteHeader(http.StatusOK)
//	err = json.NewEncoder(c.Writer).Encode(chats)
//	if err != nil {
//		fmt.Println("3")
//		return
//	}
//
//	fmt.Println("4")
//}
