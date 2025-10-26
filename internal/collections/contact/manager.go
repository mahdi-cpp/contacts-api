package contact

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/help"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
)

type Manager struct {
	userID     uuid.UUID
	onCallback OnCallback
	collection *collection_manager_memory.Manager[*Contact]
	items      map[uuid.UUID]*Contact //key is itemId
}

func NewManager(userID uuid.UUID, onCallback OnCallback, path string) (*Manager, error) {
	manager := &Manager{
		userID:     userID,
		onCallback: onCallback,
		items:      make(map[uuid.UUID]*Contact),
	}

	var err error
	manager.collection, err = collection_manager_memory.New[*Contact](path, "contacts")
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
	items, err := m.collection.ReadAll()
	if err != nil {
		return err
	}
	for _, item := range items {
		m.items[item.ID] = item
	}
	return nil
}

func (m *Manager) Create(item *Contact) (*Contact, error) {

	//err := item.Validate()
	//if err != nil {
	//	return nil, err
	//}

	// Step 2: Generate a unique ID for the new item
	itemId, err := help.GenerateUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate item ID: %w", err)
	}
	item.ID = itemId

	// Step 3: create the item in the database
	_, err = m.collection.Create(item)
	if err != nil {
		return nil, fmt.Errorf("failed to create item in database: %w", err)
	}

	err = m.load()
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *Manager) Read(itemID uuid.UUID) (*Contact, error) {

	item, err := m.collection.Read(itemID)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *Manager) ReadAll(with *SearchOptions) ([]*Contact, error) {

	all, err := m.collection.ReadAll()
	if err != nil {
		return nil, err
	}
	var items []*Contact

	filterItems := Search(all, with)

	for _, index := range filterItems {
		item, err := m.collection.Read(index.ID)
		if err != nil {
			return nil, fmt.Errorf("error reading item %s: %w", index.ID, err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (m *Manager) Update(with *UpdateOptions) error {

	if with.ID == uuid.Nil {
		return fmt.Errorf("with id is required")
	}

	item, err := m.collection.Read(with.ID)
	if err != nil {
		return fmt.Errorf("failed to read item %s: %w", with.ID, err)
	}

	Update(item, with)

	_, err = m.collection.Update(item)
	if err != nil {
		return fmt.Errorf("failed to update item %s: %w", with.ID, err)
	}

	return nil
}

func (m *Manager) Delete(itemID uuid.UUID) error {

	err := m.collection.Delete(itemID)
	if err != nil {
		fmt.Println("error deleting item")
		return err
	}

	return nil
}

func (m *Manager) Clone(path, newName string) error {

	all, err := m.collection.ReadAll()
	if err != nil {
		return err
	}

	collection, err := collection_manager_memory.New[*Contact](path, newName)

	for _, item := range all {
		fmt.Println(item.FirstName)
		_, err := collection.Copy(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) ReadJoinContacts(joins []*Join, with *SearchOptions) ([]*Contact, error) {

	var items []*Contact

	for _, join := range joins {
		item, err := m.collection.Read(join.ContactID)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	filterItems := Search(items, with)
	return filterItems, nil
}
