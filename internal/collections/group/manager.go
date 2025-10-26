package group

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/collections/contact"
	"github.com/mahdi-cpp/iris-tools/collection_manager_join"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
)

type Manager struct {
	contactManager *contact.Manager
	collection     *collection_manager_memory.Manager[*Group]
	join           *collection_manager_join.Manager[*contact.Join]
	groups         []*Group
	contacts       map[uuid.UUID][]*contact.Contact //key is groupId
}

func NewManager(contactManager *contact.Manager, path string) (*Manager, error) {

	manager := &Manager{
		contactManager: contactManager,
		contacts:       make(map[uuid.UUID][]*contact.Contact),
	}

	var err error
	manager.collection, err = collection_manager_memory.New[*Group](path, "groups")
	if err != nil {
		return nil, err
	}

	manager.join, err = collection_manager_join.New[*contact.Join](path, "group_join")
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

	m.groups = []*Group{}

	for _, item := range items {

		joinItems, err := m.join.GetByParentID(item.ID)
		if err != nil {
			continue
		}

		item.Count = len(joinItems)
		with := &contact.SearchOptions{
			Sort:      "id",
			SortOrder: "desc",
			Page:      0,
			Size:      5,
		}
		groupContacts, err := m.contactManager.ReadJoinContacts(joinItems, with)
		if err != nil {
			continue
		}

		m.groups = append(m.groups, item)
		m.contacts[item.ID] = groupContacts
	}

	return nil
}

func (m *Manager) Create(group *Group) (*Group, error) {
	item, err := m.collection.Create(group)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *Manager) Read(id uuid.UUID) (*Group, error) {
	item, err := m.collection.Read(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *Manager) ReadAll(with *SearchOptions) ([]*Group, error) {
	items, err := m.collection.ReadAll()
	if err != nil {
		return nil, err
	}

	filterItems := Search(items, with)
	return filterItems, nil
}

func (m *Manager) Update(with UpdateOptions) (*Group, error) {

	item, err := m.collection.Read(with.ID)
	if err != nil {
		return nil, err
	}

	Update(item, with)

	create, err := m.collection.Update(item)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (m *Manager) Delete(id uuid.UUID) error {
	err := m.collection.Delete(id)
	if err != nil {
		return err
	}

	RemoveGroupByID(m.groups, id)

	all, err := m.join.ReadAll()
	if err != nil {
		return err
	}

	for _, item := range all {
		if item.ParentID == id {
			err := m.join.Delete(item.GetCompositeKey())
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

	return nil
}

func (m *Manager) IsExist(id uuid.UUID) error {
	_, err := m.collection.Read(id)
	if err != nil {
		return fmt.Errorf("groups not found: %s", id)
	}

	return nil
}

//--- contact

func (m *Manager) AddContact(itemID, contactID uuid.UUID) error {

	if itemID == uuid.Nil {
		return fmt.Errorf("itemID must not be an empty string")
	}
	if contactID == uuid.Nil {
		return fmt.Errorf("contactID must not be an empty string")
	}

	j := &contact.Join{
		ParentID:  itemID,
		ContactID: contactID,
	}

	_, err := m.join.Create(j)
	if err != nil {
		return err
	}

	err = m.load()
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) DeleteContact(groupID, contactID uuid.UUID) error {

	j := &contact.Join{
		ParentID:  groupID,
		ContactID: contactID,
	}

	err := m.join.Delete(j.GetCompositeKey())
	if err != nil {
		return err
	}

	err = m.load()
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) ReadCollection(id uuid.UUID) (*Group, error) {
	item, err := m.collection.Read(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *Manager) ReadCollections(with *SearchOptions) []*contact.Collection[*Group] {

	var results []*contact.Collection[*Group]

	filterGroups := Search(m.groups, with)

	for _, group := range filterGroups {
		collection := &contact.Collection[*Group]{
			Item:     group,
			Contacts: m.contacts[group.ID],
		}
		results = append(results, collection)
	}

	return results
}

func (m *Manager) ReadCollectionContacts(with *contact.SearchOptions) ([]*contact.Contact, error) {

	groupItems, err := m.join.GetByParentID(with.ID)
	if err != nil {
		return nil, err
	}

	groupPhotos, err := m.contactManager.ReadJoinContacts(groupItems, with)
	if err != nil {
		return nil, err
	}

	return groupPhotos, nil
}

//--- events

func (m *Manager) HandleContactCreate(id uuid.UUID) {

}

func (m *Manager) HandleContactUpdate(id uuid.UUID) {

}

func (m *Manager) HandleContactDelete(id uuid.UUID) {

}

func RemoveGroupByID(groups []*Group, idToRemove uuid.UUID) {
	// ۱. پیدا کردن شاخص آلبوم
	index := -1
	for i, group := range groups {
		// از آنجایی که عناصر اشاره‌گر هستند، باید به مقدار اشاره شود (*groups) یا از طریق فیلد ID به آن دسترسی پیدا کرد.
		if group != nil && group.ID == idToRemove {
			index = i
			break // اولین مورد منطبق پیدا شد، حلقه را متوقف کنید.
		}
	}

	// اگر آلبوم پیدا نشد، اسلایس اصلی را برگردانید.
	if index == -1 {
		return
	}

	// ۲. حذف عنصر
	// این رایج‌ترین روش برای حذف عنصر در Go است که **ترتیب عناصر را حفظ می‌کند.**
	// این کار با گرفتن اسلایس قبل از شاخص، و الحاق اسلایس بعد از شاخص به آن انجام می‌شود.
	// از عملگر '...' برای بسط دادن اسلایس دوم به عنوان آرگومان‌های جداگانه به تابع append استفاده می‌شود.
	groups = append(groups[:index], groups[index+1:]...)
}
