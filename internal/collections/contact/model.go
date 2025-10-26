package contact

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (a *Join) GetRecordSize() int { return 110 }
func (a *Join) GetCompositeKey() string {
	return fmt.Sprintf("%s:%s", a.ParentID.String(), a.ContactID.String())
}

type Join struct {
	ParentID  uuid.UUID `json:"parentID"`
	ContactID uuid.UUID `json:"contactID"`
}

func (c *Contact) SetID(id uuid.UUID)       { c.ID = id }
func (c *Contact) GetID() uuid.UUID         { return c.ID }
func (c *Contact) SetCreatedAt(t time.Time) { c.CreatedAt = t }
func (c *Contact) SetUpdatedAt(t time.Time) { c.UpdatedAt = t }
func (c *Contact) GetRecordSize() int       { return 4000 }

type Contact struct {
	ID           uuid.UUID  `json:"id"`
	FirstName    string     `json:"firstName"`
	LastName     string     `json:"lastName"`
	Description  string     `json:"description"`
	Phones       []Phone    `json:"phones"`
	Emails       []Email    `json:"emails,omitempty"`
	Addresses    []Address  `json:"addresses,omitempty"`
	Profile      []Profile  `json:"profiles,omitempty"`
	Location     Location   `json:"location,omitempty"`
	Birthday     time.Time  `json:"birthday,omitempty"`
	Company      string     `json:"company"`
	Tests        []string   `json:"tests,omitempty"`
	OriginalURL  string     `json:"originalUrl"`
	ThumbnailURL string     `json:"thumbnailUrl"`
	Theme        string     `json:"theme,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
	Version      string     `json:"version"`
}

type Phone struct {
	ID    uuid.UUID `json:"id"`
	Value string    `json:"value"`
	Label string    `json:"label"`
}

type Email struct {
	ID    uuid.UUID `json:"id"`
	Value string    `json:"value"`
	Label string    `json:"label"`
}

type Address struct {
	ID    uuid.UUID `json:"id"`
	Value string    `json:"value"`
	Label string    `json:"label"`
}

type Profile struct {
	ID    uuid.UUID `json:"id"`
	Value string    `json:"value"`
	Label string    `json:"label"`
}

type Company struct {
	ID    uuid.UUID `json:"id"`
	Value string    `json:"value"`
	Label string    `json:"label"`
}

type Location struct {
	ID        uuid.UUID `json:"id"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	Address   string    `json:"address"`
}
