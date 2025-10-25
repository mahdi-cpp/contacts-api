package chat

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/collections/member"
	"github.com/mahdi-cpp/contacts-api/internal/help"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
)

type Manager struct {
	collection    *collection_manager_memory.Manager[*Chat]
	MemberManager *member.Manager
	chats         map[uuid.UUID]*Chat //key is chatID
}

func NewManager(path string, name string) (*Manager, error) {
	manager := &Manager{
		chats: make(map[uuid.UUID]*Chat),
	}

	var err error
	manager.collection, err = collection_manager_memory.New[*Chat](path, name)
	if err != nil {
		return nil, err
	}

	manager.MemberManager, err = member.NewManager(path, "members")
	if err != nil {
		return nil, err
	}

	err = manager.load()
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) load() error {
	chats, err := m.collection.ReadAll()
	if err != nil {
		return err
	}
	for _, chat := range chats {
		m.chats[chat.ID] = chat
	}
	return nil
}

func (m *Manager) CreateChat(requestChat *Chat) (*Chat, error) {

	err := requestChat.Validate()
	if err != nil {
		return nil, err
	}

	// Step 2: Generate a unique ID for the new chat
	chatID, err := help.GenerateUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate chat ID: %w", err)
	}
	requestChat.ID = chatID

	// Step 3: create the chat in the database
	_, err = m.collection.Create(requestChat)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat in database: %w", err)
	}

	err = m.load()
	if err != nil {
		return nil, err
	}

	return requestChat, nil
}

func (m *Manager) AddMember(mem *member.Member) (*member.Member, error) {
	create, err := m.MemberManager.Create(mem)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (m *Manager) ReadChat(chatID uuid.UUID) (*Chat, error) {

	chat1, err := m.collection.Read(chatID)
	if err != nil {
		return nil, err
	}

	return chat1, nil
}

func (m *Manager) ReadUserChats(userID uuid.UUID) ([]*ChatDTO, error) {

	options := &member.SearchOptions{ //first find status is member in which chats
		UserID:    userID,
		Sort:      "updatedAt",
		SortOrder: "desc",
		Page:      1,
		Size:      1000,
	}
	members, err := m.MemberManager.ReadWith(options)
	if err != nil {
		return nil, err
	}

	//fmt.Println("members: ", len(members))

	var chats []*Chat
	var dto []*ChatDTO

	for _, mem := range members {
		chat := m.chats[mem.ChatID]
		if chat != nil {
			chats = append(chats, chat)
		}
	}

	for _, chat := range chats {
		with := &member.SearchOptions{
			ChatID:    chat.ID,
			Sort:      "updatedAt",
			SortOrder: "desc",
			Page:      1,
			Size:      4,
		}
		members, err := m.MemberManager.ReadWith(with)
		if err != nil {
			return nil, err
		}

		dto = append(dto, &ChatDTO{Chat: chat, Members: members})
	}

	return dto, nil
}

func (m *Manager) UpdateChat(with *UpdateOptions) error {

	if with.ID == uuid.Nil {
		return fmt.Errorf("with id is required")
	}

	ch, err := m.collection.Read(with.ID)
	if err != nil {
		return fmt.Errorf("failed to read chat %s: %w", with.ID, err)
	}

	Update(ch, with)

	_, err = m.collection.Update(ch)
	if err != nil {
		return fmt.Errorf("failed to update chat %s: %w", with.ID, err)
	}

	return nil
}

func (m *Manager) ChatDelete(chatID uuid.UUID) error {

	err := m.collection.Delete(chatID)
	if err != nil {
		fmt.Println("error deleting chat")
		return err
	}

	return nil
}

func (m *Manager) Clone(path, newName string) error {

	all, err := m.collection.ReadAll()
	if err != nil {
		return err
	}

	collection, err := collection_manager_memory.New[*Chat](path, newName)

	for _, item := range all {
		fmt.Println(item.Title)
		_, err := collection.Copy(item)
		if err != nil {
			return err
		}
	}
	return nil
}
