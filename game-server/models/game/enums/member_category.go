package enums

type EMemberCategory int

const (
	MemberCategoryDefault EMemberCategory = iota
	MemberCategoryDeveloper
	MemberCategoryUniqueId
	MemberCategoryTrader                        EMemberCategory = 4
	MemberCategoryGroup                         EMemberCategory = 8
	MemberCategorySystem                        EMemberCategory = 16
	MemberCategoryChatModerator                 EMemberCategory = 32
	MemberCategoryChatModeratorWithPermanentBan EMemberCategory = 64
	MemberCategoryUnitTest                      EMemberCategory = 128
	MemberCategorySherpa                        EMemberCategory = 256
	MemberCategoryEmissary                      EMemberCategory = 512
)
