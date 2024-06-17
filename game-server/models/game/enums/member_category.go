package enums

type EMemberCategory int

const (
	MemberCategoryDefault EMemberCategory = iota
	MemberCategoryDeveloper
	MemberCategoryUniqueId
	MemberCategoryTrader                        = 4
	MemberCategoryGroup                         = 8
	MemberCategorySystem                        = 16
	MemberCategoryChatModerator                 = 32
	MemberCategoryChatModeratorWithPermanentBan = 64
	MemberCategoryUnitTest                      = 128
	MemberCategorySherpa                        = 256
	MemberCategoryEmissary                      = 512
)
