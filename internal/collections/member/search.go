package member

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/search"
)

type SearchOptions struct {
	ID           uuid.UUID  `json:"id,omitempty"`
	UserID       uuid.UUID  `json:"userId,omitempty"`
	ChatID       uuid.UUID  `json:"ChatId,omitempty"`
	Role         *string    `json:"role,omitempty"`        // "member", "admin", "creator"
	CustomTitle  *string    `json:"customTitle,omitempty"` // For custom admin titles
	IsActive     *bool      `json:"isActive,omitempty"`
	IsPin        *bool      `json:"isPin"`
	Theme        *string    `json:"theme"`
	Notification *time.Time `json:"notification"`

	// Date filters
	CreatedAfter  *time.Time `json:"createdAfter,omitempty"`
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`
	ActiveAfter   *time.Time `json:"activeAfter,omitempty"`

	// Pagination
	Page int `json:"page,omitempty"`
	Size int `json:"size,omitempty"`

	// Sorting
	Sort      string `json:"sort,omitempty"`      // "title", "created", "members", "lastActivity"
	SortOrder string `json:"sortOrder,omitempty"` // "asc" or "desc"
}

const MaxLimit = 1000

var LessFunks = map[string]search.LessFunction[*Member]{
	"id":        func(a, b *Member) bool { return a.ID.String() < b.ID.String() },
	"createdAt": func(a, b *Member) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"updatedAt": func(a, b *Member) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*Member] {

	fn, exists := LessFunks[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "desc" {
		return func(a, b *Member) bool { return !fn(a, b) }
	}
	return fn
}

func BuildMemberSearch(with *SearchOptions) search.Criteria[*Member] {

	return func(member *Member) bool {

		// ID filter
		if with.ID != uuid.Nil && member.ID != with.ID {
			return false
		}

		if with.UserID != uuid.Nil && member.UserID != with.UserID {
			return false
		}

		if with.ChatID != uuid.Nil && member.ChatID != with.ChatID {
			return false
		}

		if with.Role != nil && member.Role != *with.Role {
			return false
		}

		if with.CustomTitle != nil && member.CustomTitle != *with.CustomTitle {
			return false
		}

		if with.IsActive != nil && member.IsActive != *with.IsActive {
			return false
		}

		if with.IsPin != nil && member.IsPin != *with.IsPin {
			return false
		}
		if with.Theme != nil && member.Theme != *with.Theme {
			return false
		}
		if with.Notification != nil && member.Notification != *with.Notification {
			return false
		}

		// Date filters
		if with.CreatedAfter != nil && member.CreatedAt.Before(*with.CreatedAfter) {
			return false
		}
		if with.CreatedBefore != nil && member.CreatedAt.After(*with.CreatedBefore) {
			return false
		}

		return true
	}
}

func Search(items []*Member, with *SearchOptions) []*Member {

	// Build criteria
	criteria := BuildMemberSearch(with)

	// Execute search_manager
	results := search.Find(items, criteria)

	// Sort results if needed
	if with.Sort != "" {
		lessFn := GetLessFunc(with.Sort, with.SortOrder)
		if lessFn != nil {
			search.SortIndexedItems(results, lessFn)
		}
	}

	// Extract final photos
	final := make([]*Member, len(results))
	for i, item := range results {
		final[i] = item.Value
	}

	if with.Size == 0 { // if not set default is MAX_LIMIT
		with.Size = MaxLimit
	}

	// Apply pagination
	start := (with.Page - 1) * with.Size // Corrected pagination logic
	if start < 0 {
		start = 0
	}

	// Check if the start index is out of bounds. If so, return an empty slice.
	if start >= len(final) {
		return []*Member{}
	}

	end := start + with.Size
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}
