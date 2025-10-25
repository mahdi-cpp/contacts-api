package message

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/update"
)

type UpdateOptions struct {
	ID               uuid.UUID    `json:"id,omitempty"`
	Caption          *string      `json:"caption,omitempty"`
	ReplyToMessageID *string      `json:"replyToMessageId,omitempty"`
	ForwardedFrom    *ForwardInfo `json:"forwardedFrom,omitempty"`
	Entities         []Entity     `json:"entities,omitempty"`
	Reactions        []Reaction   `json:"reactions,omitempty"`
	IsEdited         *bool        `json:"isEdited,omitempty"`
	IsPinned         *bool        `json:"isPinned,omitempty"`
	IsDeleted        *bool        `json:"isDeleted,omitempty"`
	Poll             *Poll        `json:"poll,omitempty"`
	Location         *Location    `json:"location,omitempty"`
	Contact          *Contact     `json:"contact,omitempty"`
}

// Initialize updater
var metadataUpdater = update.NewUpdater[Message, UpdateOptions]()

func init() {

	// Configure scalar field updates
	metadataUpdater.AddScalarUpdater(func(a *Message, u UpdateOptions) {
		if u.Caption != nil {
			a.Caption = *u.Caption
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *Message, u UpdateOptions) {
		if u.IsEdited != nil {
			a.IsEdited = *u.IsEdited
		}
		if u.IsPinned != nil {
			a.IsPinned = *u.IsPinned
		}
		if u.IsDeleted != nil {
			a.IsDeleted = *u.IsDeleted
		}
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Message) {
		a.UpdatedAt = time.Now()
	})
}

func Update(p *Message, update *UpdateOptions) *Message {
	metadataUpdater.Apply(p, *update)
	return p
}
