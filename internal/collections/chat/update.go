package chat

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/update"
)

type UpdateOptions struct {
	ID uuid.UUID `json:"iD"`

	Type             *string `json:"type"` // "private", "group", "channel", "supergroup"
	Title            *string `json:"title"`
	Username         *string `json:"username"` // Unique identifier for public channels/groups
	Description      *string `json:"description"`
	OriginalURL      *string `json:"originalURL"` // Chat profile photo
	ThumbnailURL     *string `json:"thumbnailUrl"`
	CanSetStickerSet *bool   `json:"canSetStickerSet"` // Can set sticker set
	IsVerified       *bool   `json:"isVerified"`
	IsRestricted     *bool   `json:"isRestricted"`
	IsCreator        *bool   `json:"isCreator"`
	IsScam           *bool   `json:"isScam"`
	IsFake           *bool   `json:"isFake"`

	ActiveUsernames       *[]string `json:"users,omitempty"`                 // Full users replacement
	AddActiveUsernames    []string  `json:"AddActiveUsernames,omitempty"`    // Users to add
	RemoveActiveUsernames []string  `json:"removeActiveUsernames,omitempty"` // Users to remove

	//MembersCount        *[]Member
	//AddMembers     []Member
	//RemoveMembers  []Member
	//MembersUpdates []update.NestedFieldUpdate[Member]
}

// Key extractors for nested structs
//func memberKeyExtractor(m Member) uuid.UUID { return m.UserID }

// Initialize updater
var metadataUpdater = update.NewUpdater[Chat, UpdateOptions]()

func init() {

	// Configure scalar field updates
	metadataUpdater.AddScalarUpdater(func(a *Chat, u UpdateOptions) {
		if u.Title != nil {
			a.Title = *u.Title
		}
		if u.Type != nil {
			a.Type = *u.Type
		}
		if u.Username != nil {
			a.Username = *u.Username
		}
		if u.Description != nil {
			a.Description = *u.Description
		}
		if u.OriginalURL != nil {
			a.OriginalURL = *u.OriginalURL
		}
		if u.ThumbnailURL != nil {
			a.ThumbnailURL = *u.ThumbnailURL
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *Chat, u UpdateOptions) {

	})

	// Configure collection operations
	//metadataUpdater.AddCollectionUpdater(func(a *Chat, u UpdateOptions) {
	//	op := update.CollectionUpdateOp[string]{
	//		FullReplace: u.ActiveUsernames,
	//		Add:         u.AddActiveUsernames,
	//		Remove:      u.RemoveActiveUsernames,
	//	}
	//	a.ActiveUsernames = update.ApplyCollectionUpdate(a.ActiveUsernames, op)
	//})

	// MembersCount (ID-based updates)
	//metadataUpdater.AddNestedUpdater(func(p *Chat, u UpdateOptions) {
	//
	//	op := update.CollectionUpdateOp[Member]{
	//		FullReplace: u.MembersCount,
	//		Add:         u.AddMembers,
	//		Remove:      u.RemoveMembers,
	//	}
	//	p.MembersCount = update.ApplyCollectionUpdateByID(
	//		p.MembersCount,
	//		op,
	//		memberKeyExtractor,
	//	)
	//
	//	// Apply field-level updates to existing comments
	//	p.MembersCount = update.ApplyNestedUpdate(
	//		p.MembersCount,
	//		u.MembersUpdates,
	//		memberKeyExtractor,
	//	)
	//})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Chat) {
		a.UpdatedAt = time.Now()
	})
}

func Update(p *Chat, update *UpdateOptions) *Chat {
	metadataUpdater.Apply(p, *update)
	return p
}
