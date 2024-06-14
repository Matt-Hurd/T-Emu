package types

type ENotificationType int

const (
	Ping ENotificationType = iota
	ChannelDeleted
	TraderSupply
	GroupMatchInviteAccept
	GroupMatchInviteDecline
	GroupMatchWasRemoved
	GroupMatchInviteSend
	GroupMatchInviteCancel
	GroupMatchInviteExpired
	GroupMatchLeaderChanged
	GroupMatchUserLeave
	GroupMaxCountReached
	GroupMatchStartGame
	GroupMatchRaidSettings
	GroupMatchRaidReady
	GroupMatchRaidNotReady
	GroupMatchAbort
	WrongMajorVersion
	ChatMessageReceived
	RemovedFromFriendsList
	FriendsListNewRequest
	FriendsListRequestCanceled
	TournamentWarning
	FriendsListDecline
	FriendsListAccept
	YouWasKickedFromDialogue
	YouWereAddedToIgnoreList
	YouWereRemovedToIgnoreList
	RagfairOfferSold
	RagfairRatingChange
	RagfairNewRating
	ForceLogout
	InGameBan
	InGameUnBan
	Hideout
	TraderStanding
	ProfileLevel
	SkillPoints
	HideoutAreaLevel
	AssortmentUnlockRule
	ExamineItems
	ExamineAllItems
	TraderSalesSum
	UnlockTrader
	StashRows
	ProfileLockTimer
	MasteringSkill
	ProfileExperienceDelta
	TraderStandingDelta
	TraderSalesSumDelta
	SkillPointsDelta
	MasteringSkillDelta
	UserMatched
	UserMatchOver
	UserConfirmed
)

var ENotificationTypeNames = map[ENotificationType]string{
	Ping:                       "ping",
	ChannelDeleted:             "channel_deleted",
	TraderSupply:               "trader_supply",
	GroupMatchInviteAccept:     "groupMatchInviteAccept",
	GroupMatchInviteDecline:    "groupMatchInviteDecline",
	GroupMatchWasRemoved:       "groupMatchWasRemoved",
	GroupMatchInviteSend:       "groupMatchInviteSend",
	GroupMatchInviteCancel:     "groupMatchInviteCancel",
	GroupMatchInviteExpired:    "groupMatchInviteExpired",
	GroupMatchLeaderChanged:    "groupMatchLeaderChanged",
	GroupMatchUserLeave:        "groupMatchUserLeave",
	GroupMaxCountReached:       "groupMaxCountReached",
	GroupMatchStartGame:        "groupMatchStartGame",
	GroupMatchRaidSettings:     "groupMatchRaidSettings",
	GroupMatchRaidReady:        "groupMatchRaidReady",
	GroupMatchRaidNotReady:     "groupMatchRaidNotReady",
	GroupMatchAbort:            "groupMatchAbort",
	WrongMajorVersion:          "groupMatchUserHasBadVersion",
	ChatMessageReceived:        "new_message",
	RemovedFromFriendsList:     "youAreRemovedFromFriendList",
	FriendsListNewRequest:      "friendListNewRequest",
	FriendsListRequestCanceled: "friendListRequestCancel",
	TournamentWarning:          "tournamentWarning",
	FriendsListDecline:         "friendListRequestDecline",
	FriendsListAccept:          "friendListRequestAccept",
	YouWasKickedFromDialogue:   "groupMatchYouWasKicked",
	YouWereAddedToIgnoreList:   "youAreAddToIgnoreList",
	YouWereRemovedToIgnoreList: "youAreRemoveFromIgnoreList",
	RagfairOfferSold:           "RagfairOfferSold",
	RagfairRatingChange:        "RagfairRatingChange",
	RagfairNewRating:           "RagfairNewRating",
	ForceLogout:                "ForceLogout",
	InGameBan:                  "InGameBan",
	InGameUnBan:                "InGameUnBan",
	Hideout:                    "Hideout",
	TraderStanding:             "TraderStanding",
	ProfileLevel:               "ProfileLevel",
	SkillPoints:                "SkillPoints",
	HideoutAreaLevel:           "HideoutAreaLevel",
	AssortmentUnlockRule:       "AssortmentUnlockRule",
	ExamineItems:               "ExamineItems",
	ExamineAllItems:            "ExamineAllItems",
	TraderSalesSum:             "TraderSalesSum",
	UnlockTrader:               "UnlockTrader",
	StashRows:                  "StashRows",
	ProfileLockTimer:           "ProfileLockTimer",
	MasteringSkill:             "MasteringSkill",
	ProfileExperienceDelta:     "ProfileExperienceDelta",
	TraderStandingDelta:        "TraderStandingDelta",
	TraderSalesSumDelta:        "TraderSalesSumDelta",
	SkillPointsDelta:           "SkillPointsDelta",
	MasteringSkillDelta:        "MasteringSkillDelta",
	UserMatched:                "userMatched",
	UserMatchOver:              "userMatchOver",
	UserConfirmed:              "userConfirmed",
}

func (e ENotificationType) String() string {
	return ENotificationTypeNames[e]
}
