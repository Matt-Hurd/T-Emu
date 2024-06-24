package enums

type PlayerState byte

const (
	PlayerStateNone PlayerState = iota
	PlayerStateIdle
	PlayerStateProneIdle
	PlayerStateProneMove
	PlayerStateRun
	PlayerStateSprint
	PlayerStateJump
	PlayerStateFallDown
	PlayerStateTransition
	PlayerStateBreachDoor
	PlayerStateLoot
	PlayerStatePickup
	PlayerStateOpen
	PlayerStateClose
	PlayerStateUnlock
	PlayerStateSidestep
	PlayerStateDoorInteraction
	PlayerStateApproach
	PlayerStateProne2Stand
	PlayerStateTransit2Prone
	PlayerStatePlant
	PlayerStateStationary
	PlayerStateRoll
	PlayerStateJumpLanding
	PlayerStateClimbOver
	PlayerStateClimbUp
	PlayerStateVaultingFallDown
	PlayerStateVaultingLanding
	PlayerStateBlindFire
)
