package member

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
)

type Manager struct {
	collection *collection_manager_memory.Manager[*Member]
}

func NewManager(path string, name string) (*Manager, error) {
	manager := &Manager{}

	var err error
	manager.collection, err = collection_manager_memory.New[*Member](path, name)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) Create(mem *Member) (*Member, error) {

	memberCreated, err := m.collection.Create(mem)
	if err != nil {
		return nil, err
	}

	return memberCreated, nil
}

func (m *Manager) Read(id uuid.UUID) (*Member, error) {
	item, err := m.collection.Read(id)
	if err != nil {
		fmt.Printf("Error read collection item: %v\n", err)
		return nil, err
	}

	return item, nil
}

func (m *Manager) ReadWith(with *SearchOptions) ([]*Member, error) {

	all, err := m.collection.ReadAll()
	if err != nil {
		return nil, err
	}

	members := Search(all, with)
	return members, nil
}

func (m *Manager) Update(with *UpdateOptions) error {

	item, err := m.collection.Read(with.ID)
	if err != nil {
		return fmt.Errorf("error reading member %s: %w", with.ID, err)
	}

	Update(item, *with)

	_, err = m.collection.Update(item)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) Delete(id uuid.UUID) error {

	_, err := m.collection.Read(id)
	if err != nil {
		return err
	}

	err = m.collection.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete metadata: %w", err)
	}

	return nil
}

func (m *Manager) ReadByIds(ids []uuid.UUID) ([]*Member, error) {

	members := make([]*Member, 0, len(ids))

	for _, id := range ids {
		read, err := m.collection.Read(id)
		if err != nil {
			return nil, fmt.Errorf("error reading message %s: %w", id, err)
		}
		members = append(members, read)
	}

	return members, nil
}

func (m *Manager) Copy(path, newName string) error {
	all, err := m.collection.ReadAll()
	if err != nil {
		return err
	}

	co, err := collection_manager_memory.NewWithRecordSize[*Member](path, newName, 600)

	for _, item := range all {
		_, err := co.Copy(item)
		if err != nil {
			return err
		}
	}
	return nil
}
