package group

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/update"
)

type UpdateOptions struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title,omitempty"`
	Subtitle string    `json:"subtitle,omitempty"`
	Type     string    `json:"type,omitempty"`
}

// Initialize updater
var metadataUpdater = update.NewUpdater[Group, UpdateOptions]()

func init() {

	metadataUpdater.AddScalarUpdater(func(a *Group, u UpdateOptions) {
		if u.Title != "" {
			a.Title = u.Title
		}
		if u.Subtitle != "" {
			a.Subtitle = u.Subtitle
		}
		if u.Type != "" {
			a.Type = u.Type
		}
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Group) {
		a.UpdatedAt = time.Now()
	})

}

func Update(item *Group, update UpdateOptions) *Group {
	metadataUpdater.Apply(item, update)
	return item
}
