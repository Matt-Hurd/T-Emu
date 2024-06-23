package enums

type VaultingStrategy byte

const (
	VaultingStrategyNone VaultingStrategy = iota
	Vault
	Climb
)
