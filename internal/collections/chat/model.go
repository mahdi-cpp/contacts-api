package chat

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/collections/member"
	"github.com/mahdi-cpp/contacts-api/internal/collections/status"
)

func (a *Join) GetRecordSize() int { return 110 }
func (a *Join) GetCompositeKey() string {
	return fmt.Sprintf("%s:%s", a.ParentID.String(), a.PhotoID.String())
}

type Join struct {
	ParentID uuid.UUID `json:"parentID"`
	PhotoID  uuid.UUID `json:"photoId"`
}

func (c *Chat) SetID(id uuid.UUID)       { c.ID = id }
func (c *Chat) GetID() uuid.UUID         { return c.ID }
func (c *Chat) SetCreatedAt(t time.Time) { c.CreatedAt = t }
func (c *Chat) SetUpdatedAt(t time.Time) { c.UpdatedAt = t }
func (c *Chat) GetRecordSize() int       { return 4000 }

type Chat struct {
	ID           uuid.UUID `json:"id"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	Username     string    `json:"username,omitempty"`
	Description  string    `json:"description,omitempty"`
	OriginalURL  string    `json:"originalUrl"`
	ThumbnailURL string    `json:"thumbnailUrl"`

	MembersCount int  `json:"membersCount"`
	IsVerified   bool `json:"isVerified"`

	Permissions Permissions `json:"permissions"`

	InviteLink   string    `json:"inviteLink,omitempty"`
	LinkedChatID string    `json:"linkedChatId,omitempty"` // Changed type to string for consistency
	Location     *Location `json:"location,omitempty"`

	AvailableReactions []string `json:"availableReactions,omitempty"`
	SlowModeDelay      int      `json:"slowModeDelay,omitempty"`

	Theme string `json:"theme,omitempty"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Version   string     `json:"version"`
}

type Permissions struct {
	CanSendMessages      bool `json:"canSendMessages"`
	CanSendMedia         bool `json:"canSendMedia"`
	CanSendPolls         bool `json:"canSendPolls"`
	CanSendOtherMessages bool `json:"canSendOtherMessages"`

	CanAddWebPagePreviews bool `json:"canAddWebPagePreviews"`
	CanChangeInfo         bool `json:"canChangeInfo"`
	CanInviteUsers        bool `json:"canInviteUsers"`
	CanPinMessages        bool `json:"canPinMessages"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Address   string  `json:"address"`
}

type MessagePreview struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Type      string    `json:"type"`
	AuthorID  string    `json:"authorId"` // Changed name and type for consistency
	Timestamp time.Time `json:"timestamp"`
}

type ChatDTO struct {
	Chat    *Chat            `json:"chat"`
	Members []*member.Member `json:"members"`
	Status  status.Status    `json:"status"`
}
