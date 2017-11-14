// Code generated by "stringer -type EventType -output event_types_string.go"; DO NOT EDIT.

package rtm

import "fmt"

const _EventType_name = "InvalidEventTypeClientConnectingEventTypeClientDisconnectedEventTypeAccountsChangedTypeBotAddedTypeBotChangedTypeChannelArchiveTypeChannelCreatedTypeChannelDeletedTypeChannelHistoryChangedTypeChannelJoinedTypeChannelLeftTypeChannelMarkedTypeChannelRenameTypeChannelUnarchiveTypeCommandsChangedTypeDndUpdatedTypeDndUpdatedUserTypeEmailDomainChangedTypeEmojiChangedTypeErrorTypeFileChangeTypeFileCommentAddedTypeFileCommentDeletedTypeFileCommentEditedTypeFileCreatedTypeFileDeletedTypeFilePublicTypeFileSharedTypeFileUnsharedTypeGoodbyeTypeGroupArchiveTypeGroupCloseTypeGroupHistoryChangedTypeGroupJoinedTypeGroupLeftTypeGroupMarkedTypeGroupOpenTypeGroupRenameTypeGroupUnarchiveTypeHelloTypeImCloseTypeImCreatedTypeImHistoryChangedTypeImMarkedTypeImOpenTypeManualPresenceChangeTypeMessageTypePinAddedTypePinRemovedTypePrefChangeTypePresenceChangeTypePongTypeReactionAddedTypeReactionRemovedTypeReconnectURLTypeStarAddedTypeStarRemovedTypeSubteamCreatedTypeSubteamSelfAddedTypeSubteamSelfRemovedTypeSubteamUpdatedTypeTeamDomainChangeTypeTeamJoinTypeTeamMigrationStartedTypeTeamPlanChangeTypeTeamPrefChangeTypeTeamProfileChangeTypeTeamProfileDeleteTypeTeamProfileReorderTypeTeamRenameTypeUserChangeTypeUserTypingType"

var _EventType_index = [...]uint16{0, 16, 41, 68, 87, 99, 113, 131, 149, 167, 192, 209, 224, 241, 258, 278, 297, 311, 329, 351, 367, 376, 390, 410, 432, 453, 468, 483, 497, 511, 527, 538, 554, 568, 591, 606, 619, 634, 647, 662, 680, 689, 700, 713, 733, 745, 755, 779, 790, 802, 816, 830, 848, 856, 873, 892, 908, 921, 936, 954, 974, 996, 1014, 1034, 1046, 1070, 1088, 1106, 1127, 1148, 1170, 1184, 1198, 1212}

func (i EventType) String() string {
	if i < 0 || i >= EventType(len(_EventType_index)-1) {
		return fmt.Sprintf("EventType(%d)", i)
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}
