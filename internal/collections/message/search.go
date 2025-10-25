package message

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/search"
)

const MaxLimit = 1000

type SearchOptions struct {
	ID          uuid.UUID `form:"id"`
	UserID      uuid.UUID `form:"userId"`
	Type        *string   `json:"type"`
	IsEdited    *bool     `form:"isEdited"`
	IsPinned    *bool     `form:"isPinned"`
	IsDeleted   *bool     `form:"isDeleted"`
	MediaUnread *bool     `json:"mediaUnread" `

	// Date filters
	CreatedAfter  *time.Time `form:"createdAfter"`
	CreatedBefore *time.Time `form:"createdBefore"`
	ActiveAfter   *time.Time `form:"activeAfter"`

	// Sorting
	Sort      string `form:"sort,omitempty"`
	SortOrder string `form:"sortOrder,omitempty"`

	// Pagination
	Page int `form:"page,omitempty"`
	Size int `form:"size,omitempty"`
}

var LessFunks = map[string]search.LessFunction[*Index]{
	"id":        func(a, b *Index) bool { return a.ID.String() < b.ID.String() },
	"createdAt": func(a, b *Index) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"updatedAt": func(a, b *Index) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*Index] {

	fn, exists := LessFunks[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "desc" {
		return func(a, b *Index) bool { return !fn(a, b) }
	}
	return fn
}

func BuildMessageCriteria(with *SearchOptions) search.Criteria[*Index] {

	return func(c *Index) bool {

		if with.Type != nil && c.Type != *with.Type {
			return false
		}

		if with.ID != uuid.Nil && c.ID != with.ID {
			return false
		}
		if with.UserID != uuid.Nil && c.UserID != with.UserID {
			return false
		}

		// Boolean flags
		if with.IsEdited != nil && c.IsEdited != *with.IsEdited {
			return false
		}
		if with.IsPinned != nil && c.IsPinned != *with.IsPinned {
			return false
		}
		if with.IsDeleted != nil && c.IsDeleted != *with.IsDeleted {
			return false
		}
		if with.MediaUnread != nil && c.MediaUnread != *with.MediaUnread {
			return false
		}

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

func Search(indexes []*Index, with *SearchOptions) []*Index {

	// Build criteria
	criteria := BuildMessageCriteria(with)

	// Execute search_manager
	results := search.Find(indexes, criteria)

	// Sort results if needed
	if with.Sort != "" {
		lessFn := GetLessFunc(with.Sort, with.SortOrder)
		if lessFn != nil {
			search.SortIndexedItems(results, lessFn)
		}
	}

	// Extract final assets
	final := make([]*Index, len(results))
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
		return []*Index{}
	}

	end := start + with.Size
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}
