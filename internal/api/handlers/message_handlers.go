package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mahdi-cpp/contacts-api/internal/application"
	"github.com/mahdi-cpp/iris-tools/mygin"
)

type MessageHandler struct {
	appManager *application.AppManager
}

func NewMessageHandler(appManager *application.AppManager) *MessageHandler {
	return &MessageHandler{
		appManager: appManager,
	}
}

func (h *MessageHandler) Create(c *mygin.Context) {

	//fmt.Println("ChatHandler create")

	//var request *message.Message
	//err := json.NewDecoder(c.Request.Body).Decode(&request)
	//if err != nil {
	//	http.Error(c.Writer, "Invalid request body", http.StatusBadRequest)
	//	return
	//}

	//fmt.Println(request.ChatID)
	//newMessage, err := h.appManager.Create(request)
	//if err != nil {
	//	// اینجا می‌توانید خطای مناسب را به کاربر برگردانید
	//	http.Error(c.Writer, "Failed to create message", http.StatusInternalServerError)
	//	return
	//}

	//c.Writer.Header().Set("Content-Type", "application/json")
	//c.Writer.WriteHeader(http.StatusCreated)
	//err = json.NewEncoder(c.Writer).Encode(newMessage)
	//if err != nil {
	//	return
	//}
}

func (h *MessageHandler) Read(c *mygin.Context) {

	chatID := c.GetQuery("chatId")
	fmt.Println("chatID: ", chatID)

	limit := c.GetQuery("limit")
	fmt.Println("limit: ", limit)

	//var request message.SearchOptions
	//err := json.NewDecoder(c.Request.Body).Decode(&request)
	//if err != nil {
	//	http.Error(c.Writer, "Invalid request body", http.StatusBadRequest)
	//	return
	//}

	//if err := c.ShouldBindQuery(&request); err != nil {
	//	fmt.Println(err)
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	//if request.MessageID == uuid.Nil { //read all messages with Message SearchOptions
	//	h.readAllMessage(w, r, &request)
	//} else if request.MessageID != uuid.Nil {
	//	h.readSingleMessage(w, r, request.ChatID, request.MessageID)
	//}
}

func (h *MessageHandler) Read2(c *mygin.Context) {

	// Get query parameters
	chatID := c.GetQuery("chatId")
	limitStr := c.GetQuery("limit")
	limit, err := c.GetQueryInt("limit")
	if err != nil {

	}

	if limit == 23 {
		fmt.Println("limit is 23")
	}

	// Convert limit to integer with a default value
	//limit := 10 // default value
	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil {
			limit = parsedLimit
		}
	}

	// Or use the helper method
	// limit := c.GetQueryIntDefault("limit", 10)

	// Your business logic here
	c.JSON(http.StatusOK, mygin.H{
		"chatId": chatID,
		"limit":  limit,
		"messages": []mygin.H{
			{"id": 1, "text": "Hello"},
			{"id": 2, "text": "World"},
		},
	})
}

//func (h *ChatHandler) readAllMessage(c *rest.Context, options *message.SearchOptions) {
//	fmt.Println("readAllMessage")
//
//	selectedMessages, err := h.appManager.ReadMessages(options)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, selectedMessages)
//}
//
//func (h *ChatHandler) readSingleMessage(w http.ResponseWriter, r *http.Request, chatID, messageId uuid.UUID) {
//
//	fmt.Println("readSingleMessage", chatID)
//	chatManager, err := h.appManager.GetChatManager(chatID)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//		return
//	}
//
//	readMessage, err := chatManager.ReadMessage(messageId)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, readMessage)
//}
