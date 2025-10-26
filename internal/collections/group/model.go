package group

import (
	"time"

	"github.com/google/uuid"
)

func (a *Group) GetID() uuid.UUID         { return a.ID }
func (a *Group) SetID(id uuid.UUID)       { a.ID = id }
func (a *Group) SetCreatedAt(t time.Time) { a.CreatedAt = t }
func (a *Group) SetUpdatedAt(t time.Time) { a.UpdatedAt = t }
func (a *Group) GetRecordSize() int       { return 2048 }

type Group struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Type      string    `json:"type"`
	Number    int       `json:"number"`
	IsHidden  bool      `json:"isHidden"`
	Count     int       `json:"count"`
	LastSeen  time.Time `json:"lastSeen"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Version   string    `json:"version"`
}
