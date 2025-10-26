package contact

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/search"
)

type SearchOptions struct {
	ID      uuid.UUID `form:"id,omitempty"`
	UserID  uuid.UUID `form:"userId,omitempty"`
	GroupID uuid.UUID `form:"groupId,omitempty"`

	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	Description *string `json:"description"`
	Company     *string `json:"company"`

	// Date filters
	CreatedAfter  *time.Time `form:"createdAfter,omitempty"`
	CreatedBefore *time.Time `form:"createdBefore,omitempty"`
	ActiveAfter   *time.Time `form:"activeAfter,omitempty"`

	// Sorting
	Sort      string `json:"sort,omitempty"`      // "title", "created", "MemberManager", "lastActivity"
	SortOrder string `json:"sortOrder,omitempty"` // "asc" or "desc"

	// Pagination
	Page int `json:"page,omitempty"`
	Size int `json:"size,omitempty"`
}

const MaxLimit = 1000

var LessFunks = map[string]search.LessFunction[*Contact]{
	"id":        func(a, b *Contact) bool { return a.ID.String() < b.ID.String() },
	"createdAt": func(a, b *Contact) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"updatedAt": func(a, b *Contact) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*Contact] {

	fn, exists := LessFunks[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "desc" {
		return func(a, b *Contact) bool { return !fn(a, b) }
	}
	return fn
}

func BuildChatCriteria(with *SearchOptions) search.Criteria[*Contact] {

	return func(c *Contact) bool {

		// ID filter
		if with.ID != uuid.Nil && c.ID != with.ID {
			return false
		}
		if with.FirstName != nil && c.FirstName != *with.FirstName {
			return false
		}
		if with.LastName != nil && c.LastName != *with.LastName {
			return false
		}
		if with.Description != nil && c.Description != *with.Description {
			return false
		}

		// Boolean flags
		//if with.IsVerified != nil && c.IsVerified != *with.IsVerified {
		//	return false
		//}

		// Date filters
		if with.CreatedAfter != nil && c.CreatedAt.Before(*with.CreatedAfter) {
			return false
		}
		if with.CreatedBefore != nil && c.CreatedAt.After(*with.CreatedBefore) {
			return false
		}

		return true
	}
}

func Search(chats []*Contact, with *SearchOptions) []*Contact {

	// Build criteria
	criteria := BuildChatCriteria(with)

	// Execute search_manager
	results := search.Find(chats, criteria)

	// Sort results if needed
	if with.Sort != "" {
		lessFn := GetLessFunc(with.Sort, with.SortOrder)
		if lessFn != nil {
			search.SortIndexedItems(results, lessFn)
		}
	}

	// Extract final assets
	final := make([]*Contact, len(results))
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
		return []*Contact{}
	}

	end := start + with.Size
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}
