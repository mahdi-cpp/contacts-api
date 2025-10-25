package application

import (
	"path/filepath"
	"sync"

	"github.com/mahdi-cpp/contacts-api/internal/collections/chat"
	"github.com/mahdi-cpp/contacts-api/internal/config"
)

type AppManager struct {
	mu          sync.RWMutex
	chatManager *chat.Manager
	createChat  chan *chat.Chat
}

func New() (*AppManager, error) {

	manager := &AppManager{
		createChat: make(chan *chat.Chat, 100),
	}

	var err error
	var chatsDirectory = filepath.Join(config.RootDir, "metadata")
	manager.chatManager, err = chat.NewManager(chatsDirectory, "chats")
	if err != nil {
		panic(err)
	}

	return manager, nil
}

func (m *AppManager) GetChatManager() (*chat.Manager, error) {

	//messageManager, ok := m.messageManagers[chatID]
	//if ok {
	//	return messageManager, nil
	//}

	//chat1, err := m.ChatCollectionManager.Read(chatID)
	//if err != nil {
	//	fmt.Println("chat not found in cash")
	//	return nil, err
	//}

	//messageManager, err := message.New(chatID, "")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//m.messageManagers[chatID] = messageManager // add to cash

	return m.chatManager, nil
}
