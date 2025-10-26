package contact

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/update"
)

type UpdateOptions struct {
	ID          uuid.UUID `json:"id"`
	FirstName   *string   `json:"firstName"`
	LastName    *string   `json:"lastName"`
	Description *string   `json:"description"`

	Emails       []Email    `json:"emails,omitempty"`
	Addresses    []Address  `json:"addresses,omitempty"`
	Profile      []Profile  `json:"profiles,omitempty"`
	Location     Location   `json:"location,omitempty"`
	Birthday     *time.Time `json:"birthday,omitempty"`
	Company      *string    `json:"company"`
	OriginalURL  *string    `json:"originalUrl"`
	ThumbnailURL *string    `json:"thumbnailUrl"`
	Theme        *string    `json:"theme,omitempty"`

	Tests       *[]string `json:"tests,omitempty"`       // Full tests replacement
	AddTests    []string  `json:"addTests,omitempty"`    // Tests to add
	RemoveTests []string  `json:"removeTests,omitempty"` // Tests to remove

	Phones        *[]Phone `json:"phones,omitempty"`
	AddPhones     []Phone
	RemovePhones  []Phone
	PhonesUpdates []update.NestedFieldUpdate[Phone]
}

// Key extractors for nested structs
func phoneKeyExtractor(p Phone) uuid.UUID { return p.ID }

// Initialize updater
var metadataUpdater = update.NewUpdater[Contact, UpdateOptions]()

func init() {

	// Configure scalar field updates
	metadataUpdater.AddScalarUpdater(func(a *Contact, u UpdateOptions) {
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

	metadataUpdater.AddScalarUpdater(func(a *Contact, u UpdateOptions) {

	})

	// Configure collection operations
	metadataUpdater.AddCollectionUpdater(func(a *Contact, u UpdateOptions) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Tests,
			Add:         u.AddTests,
			Remove:      u.RemoveTests,
		}
		a.Tests = update.ApplyCollectionUpdate(a.Tests, op)
	})

	//Phones (ID-based updates)
	metadataUpdater.AddNestedUpdater(func(p *Contact, u UpdateOptions) {

		op := update.CollectionUpdateOp[Phone]{
			FullReplace: u.Phones,
			Add:         u.AddPhones,
			Remove:      u.RemovePhones,
		}
		p.Phones = update.ApplyCollectionUpdateByID(
			p.Phones,
			op,
			phoneKeyExtractor,
		)

		// Apply field-level updates to existing comments
		p.Phones = update.ApplyNestedUpdate(
			p.Phones,
			u.PhonesUpdates,
			phoneKeyExtractor,
		)
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Contact) {
		a.UpdatedAt = time.Now()
	})
}

func Update(p *Contact, update *UpdateOptions) *Contact {
	metadataUpdater.Apply(p, *update)
	return p
}
