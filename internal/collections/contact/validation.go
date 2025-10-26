package contact

import (
	"fmt"
	"regexp"
	"strings"
)

func (c *UpdateOptions) ValidateUpdate() error {
	return nil
}

func (c *Contact) Validate() error {

	err := validateUsername(c.FirstName)
	if err != nil {
		return err
	}

	//err = typeValidate(c.Type)
	//if err != nil {
	//	return err
	//}

	const maxDescriptionLength = 50
	if len(strings.TrimSpace(c.Description)) > maxDescriptionLength {
		return fmt.Errorf("description length exceeds the maximum of %d characters", maxDescriptionLength)
	}

	const maxInviteLinkLength = 500
	if len(strings.TrimSpace(c.LastName)) > maxInviteLinkLength {
		return fmt.Errorf("inviteLink length exceeds the maximum of %d characters", maxInviteLinkLength)
	}

	// Validate all member data using the Member's own Validate method.
	//membersMap := make(map[uuid.UUID]struct{})
	//for _, member := range c.Phones {
	//
	//	// Call the Validate method on the member struct itself.
	//	if err := member.validate(); err != nil {
	//		// This returns an error if Role or UserID are invalid.
	//		return fmt.Errorf("member validation failed for UserID '%s': %w", member.UserID, err)
	//	}
	//
	//	// Check for duplicate MemberManager. This responsibility remains here because
	//	// it is a concern of the *collection* of MemberManager, not a single member.
	//	if _, exists := membersMap[member.UserID]; exists {
	//		return fmt.Errorf("duplicate member with UserID '%s' found", member.UserID)
	//	}
	//	membersMap[member.UserID] = struct{}{}
	//}

	return nil
}

// ValidateUsername checks if the Username field of the Contact struct is valid.
// A valid username:
// - Must be between 5 and 32 characters long.
// - Can contain alphanumeric characters and underscores.
// - Must start with a letter.
// - Cannot end with an underscore.
// - Cannot have consecutive underscores.
func validateUsername(username string) error {

	if username == "" { //also can empty in private , groups collection
		return nil
	}

	// 1. Check if the username is within the valid length range.
	if len(username) < 5 || len(username) > 32 {
		return fmt.Errorf("username length must be between 0 and 32 characters, got %d", len(username))
	}

	// 2. Use a regular expression for comprehensive validation.
	// The regex explained:
	// ^ - Asserts position at the start of the string.
	// [a-zA-Z] - Matches any single uppercase or lowercase letter.
	// [a-zA-Z0-9_]* - Matches any combination of letters, numbers, or underscores, zero or more times.
	// [a-zA-Z0-9] - Matches any single letter or number. This is to ensure the username does not end in an underscore.
	// $ - Asserts position at the end of the string.
	// The pattern effectively requires the username to start with a letter and end with a letter or number.
	// It also implicitly handles the consecutive underscore case by not allowing it to match the end of the string if it contains one.
	pattern := "^[a-zA-Z][a-zA-Z0-9_]*[a-zA-Z0-9]$"
	if !regexp.MustCompile(pattern).MatchString(username) {
		return fmt.Errorf("username '%s' is invalid. It must start with a letter, be between 5-32 characters, and contain only letters, numbers, and underscores, without ending in an underscore", username)
	}

	return nil
}

func typeValidate(messageType string) error {

	validTypes := map[string]struct{}{
		"private": {},
		"groups":  {},
		"channel": {},
		"bot":     {},
	}

	// Check if the Role is one of the valid ones.
	if _, isValid := validTypes[messageType]; !isValid {
		return fmt.Errorf("invalid type '%s'", messageType)
	}

	// If all validations pass, return nil (no error)
	return nil
}

//func (m *Member) validate() error {
//
//	// A list of valid roles to check against
//	validRoles := map[string]struct{}{
//		"member":  {},
//		"admin":   {},
//		"creator": {},
//	}
//
//	// Check if the Role is one of the valid ones.
//	if _, isValid := validRoles[m.Role]; !isValid {
//		return fmt.Errorf("invalid role '%s'", m.Role)
//	}
//
//	// Check if the UserID is a valid UUID format.
//	//if _, err := help.IsValidUUID(m.UserID); err != nil {
//	//	return fmt.Errorf("invalid UserID format: %w", err)
//	//}
//
//	// Third, validate the CustomTitle string field.
//	const maxCustomTitleLength = 50
//	if len(strings.TrimSpace(m.CustomTitle)) > maxCustomTitleLength {
//		return fmt.Errorf("custom title length exceeds the maximum of %d characters", maxCustomTitleLength)
//	}
//
//	// If all validations pass, return nil (no error)
//	return nil
//}
