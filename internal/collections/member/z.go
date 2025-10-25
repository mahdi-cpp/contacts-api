package member

//
//// Chat-specific search functions
//// ---------------------------------------------------------------------
//
//// HasMemberWith creates a criteria that checks if a chat has at least one member
//// matching the provided member criteria
//func HasMemberWith(memberCriteria search.Criteria[*Member]) search.Criteria[*Chat] {
//	return func(chat *Chat) bool {
//		for _, member := range chat.MembersCount {
//			if memberCriteria(&member) {
//				return true
//			}
//		}
//		return false
//	}
//}
//
//// Member-specific criteria functions
//// ---------------------------------------------------------------------
//
//// MemberWithUserID checks if a member has a specific status ID
//func MemberWithUserID(userID uuid.UUID) search.Criteria[*Member] {
//	return func(member *Member) bool {
//		return member.UserID == userID
//	}
//}
//
//// MemberWithRole checks if a member has a specific role
//func MemberWithRole(role string) search.Criteria[*Member] {
//	return func(member *Member) bool {
//		return member.Role == role
//	}
//}
//
//// MemberWithCustomTitle checks if a member's custom title contains the query
//func MemberWithCustomTitle(query string) search.Criteria[*Member] {
//	return func(member *Member) bool {
//		return search.StringContains(member.CustomTitle, query)
//	}
//}
//
//// MemberIsActive checks if a member is active
//func MemberIsActive() search.Criteria[*Member] {
//	return func(member *Member) bool {
//		return member.IsActive
//	}
//}
//
//// MemberJoinedAfter checks if a member joined after a specific time
//func MemberJoinedAfter(time time.Time) search.Criteria[*Member] {
//	return func(member *Member) bool {
//		return member.JoinedAt.After(time)
//	}
//}
//
//// MemberJoinedBefore checks if a member joined before a specific time
//func MemberJoinedBefore(time time.Time) search.Criteria[*Member] {
//	return func(member *Member) bool {
//		return member.JoinedAt.Before(time)
//	}
//}
//
//// ActiveAdminsJoinedAfter finds active admins who joined after a specific time
//func ActiveAdminsJoinedAfter(time time.Time) search.Criteria[*Member] {
//	return func(member *Member) bool {
//		return member.Role == "admin" && member.IsActive && member.JoinedAt.After(time)
//	}
//}
//
//// MembersWithTitlePattern finds members with a specific title pattern
//func MembersWithTitlePattern(pattern string) search.Criteria[*Member] {
//	return func(member *Member) bool {
//		return strings.Contains(strings.ToLower(member.CustomTitle),
//			strings.ToLower(pattern))
//	}
//}
//
//// CountMembers returns the number of members matching the criteria in a chat
//func CountMembers(chat *Chat, criteria search.Criteria[*Member]) int {
//	count := 0
//	for _, member := range chat.MembersCount {
//		if criteria(&member) {
//			count++
//		}
//	}
//	return count
//}
//
//// GetMatchingMembers returns all members in a chat that match the criteria
//func GetMatchingMembers(chat Chat, criteria search.Criteria[*Member]) []Member {
//	var matches []Member
//	for _, member := range chat.MembersCount {
//		if criteria(&member) {
//			matches = append(matches, member)
//		}
//	}
//	return matches
//}
//
//// Sorting functions for members
//// ---------------------------------------------------------------------
//
//// SortByRole sorts members by role (creator > admin > member)
//func SortByRole(members []Member) {
//	rolePriority := map[string]int{
//		"creator": 0,
//		"admin":   1,
//		"member":  2,
//	}
//
//	sort.Slice(members, func(i, j int) bool {
//		return rolePriority[members[i].Role] < rolePriority[members[j].Role]
//	})
//}
//
//// SortByJoinedAt sorts members by join date (newest first by default)
//func SortByJoinedAt(members []Member, ascending bool) {
//	sort.Slice(members, func(i, j int) bool {
//		if ascending {
//			return members[i].JoinedAt.Before(members[j].JoinedAt)
//		}
//		return members[i].JoinedAt.After(members[j].JoinedAt)
//	})
//}
//
//// SortByLastActive sorts members by last activity time (most recent first by default)
//func SortByLastActive(members []Member, ascending bool) {
//	sort.Slice(members, func(i, j int) bool {
//		if ascending {
//			return members[i].LastActive.Before(members[j].LastActive)
//		}
//		return members[i].LastActive.After(members[j].LastActive)
//	})
//}
//
//// SortByActivityStatus sorts members (active first, then inactive)
//func SortByActivityStatus(members []Member) {
//	sort.Slice(members, func(i, j int) bool {
//		// Active members first
//		if members[i].IsActive && !members[j].IsActive {
//			return true
//		}
//		if !members[i].IsActive && members[j].IsActive {
//			return false
//		}
//		// If both have same status, sort by last active
//		return members[i].LastActive.After(members[j].LastActive)
//	})
//}
//
//// SortByRoleThenJoinDate Multi-level sorting: First by role, then by join date
//func SortByRoleThenJoinDate(members []Member) {
//	sort.Slice(members, func(i, j int) bool {
//		// First, sort by role priority
//		rolePriority := map[string]int{
//			"creator": 0,
//			"admin":   1,
//			"member":  2,
//		}
//
//		if rolePriority[members[i].Role] != rolePriority[members[j].Role] {
//			return rolePriority[members[i].Role] < rolePriority[members[j].Role]
//		}
//
//		// If same role, sort by join date (newest first)
//		return members[i].JoinedAt.After(members[j].JoinedAt)
//	})
//}
