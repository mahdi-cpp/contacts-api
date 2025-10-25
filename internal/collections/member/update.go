package member

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/update"
)

type UpdateOptions struct {
	ID          uuid.UUID `json:"id"`
	Role        *string   `json:"role"`        // "member", "admin", "creator"
	CustomTitle *string   `json:"customTitle"` // For custom admin titles
	IsActive    *bool     `json:"isActive"`

	IsPin        *bool      `json:"isPin"`
	Theme        *string    `json:"theme"`
	Notification *time.Time `json:"notification"`
}

// Initialize updater
var metadataUpdater = update.NewUpdater[Member, UpdateOptions]()

func init() {

	metadataUpdater.AddScalarUpdater(func(member *Member, u UpdateOptions) {
		if u.Role != nil {
			member.Role = *u.Role
		}
		if u.CustomTitle != nil {
			member.CustomTitle = *u.CustomTitle
		}
		if u.IsActive != nil {
			member.IsActive = *u.IsActive
		}

		if u.IsPin != nil {
			member.IsPin = *u.IsPin
		}
		if u.Theme != nil {
			member.Theme = *u.Theme
		}
		if u.Notification != nil {
			member.Notification = *u.Notification
		}
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Member) {
		a.UpdatedAt = time.Now()
	})
}

func Update(item *Member, update UpdateOptions) *Member {
	metadataUpdater.Apply(item, update)
	return item
}
