package member

import (
	"time"

	"github.com/google/uuid"
)

func (a *Member) GetID() uuid.UUID         { return a.ID }
func (a *Member) SetID(id uuid.UUID)       { a.ID = id }
func (a *Member) SetCreatedAt(t time.Time) { a.CreatedAt = t }
func (a *Member) SetUpdatedAt(t time.Time) { a.UpdatedAt = t }
func (a *Member) GetRecordSize() int       { return 600 }

type Member struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userId"`
	ChatID      uuid.UUID `json:"ChatId"`
	Role        string    `json:"role"`        // "creator", "member", "admin"
	CustomTitle string    `json:"customTitle"` // For custom admin titles
	IsActive    bool      `json:"isActive"`

	IsPin        bool      `json:"isPin"`
	Theme        string    `json:"theme"`
	Notification time.Time `json:"notification"`

	CreatedAt time.Time `json:"createdAt"` // joined Date
	UpdatedAt time.Time `json:"updatedAt"` // Added for sorting by activity
	DeletedAt time.Time `json:"deletedAt"`
}
