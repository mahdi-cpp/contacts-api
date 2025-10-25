package message

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_index"
)

type Manager struct {
	mu         sync.RWMutex
	userID     uuid.UUID
	collection *collection_manager_index.Manager[*Message, *Index]
}

func New(userID uuid.UUID, path string) (*Manager, error) {
	manager := &Manager{
		userID: userID,
	}

	var err error
	manager.collection, err = collection_manager_index.New[*Message, *Index](path)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

// Create adds a new message to the chat. No context is passed here.
func (m *Manager) Create(msg *Message) error {
	_, err := m.collection.Create(msg)
	if err != nil {
		return err
	}
	return nil
}

// Read retrieves a message by its ID.
func (m *Manager) Read(id uuid.UUID) (*Message, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("cannot read message with an empty ID")
	}
	msg, err := m.collection.Read(id)
	if err != nil {
		return nil, fmt.Errorf("error reading message %s: %w", id, err)
	}
	return msg, nil
}

// ReadAll retrieves all messages in this chat.
func (m *Manager) ReadAll(with *SearchOptions) ([]*Message, error) {

	all := m.collection.GetAllIndexes()
	var messages []*Message

	filterIndexes := Search(all, with)

	for _, index := range filterIndexes {
		read, err := m.collection.Read(index.ID)
		if err != nil {
			return nil, fmt.Errorf("error reading message %s: %w", index.ID, err)
		}
		messages = append(messages, read)
	}

	return messages, nil
}

// Update updates a
func (m *Manager) Update(with *UpdateOptions) (*Message, error) {

	if with.ID == uuid.Nil {
		return nil, fmt.Errorf("cannot update message without an ID")
	}

	msg, err := m.collection.Read(with.ID)
	if err != nil {
		return nil, err
	}
	Update(msg, with)
	msg, err = m.collection.Update(msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// Delete deletes a
func (m *Manager) Delete(id uuid.UUID) error {
	if id == uuid.Nil {
		return fmt.Errorf("cannot delete message without an ID")
	}

	err := m.collection.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
