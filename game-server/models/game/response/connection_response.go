package response

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
	"log"
	"path/filepath"
	"time"
)

type PacketConnection struct {
	EncryptionEnabled        bool
	DecryptionEnabled        bool
	GameDateTime             core.GameDateTime
	ResourceKeyArrayZipped   []byte
	CustomizationArrayZipped []byte
	WeatherArrayZipped       []byte
	Season                   byte
	CanRestart               bool
	EmemberCategory          EMemberCategory
	FixedDeltaTime           float32
	InteractablesZipped      []byte
	SessionId                []byte
	Bounds                   core.Bounds
	SCPort                   uint16
	EnetLogsLevel            ENetLogsLevel
	GitVersion               core.GitVersion
	SpeedLimitEnabled        bool
	Config                   core.ConnectionConfig
	VoipSettings             *core.VoipSettings
}

type EMemberCategory int32
type ENetLogsLevel byte

func (rp *PacketConnection) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.ReadBool(buffer, &rp.EncryptionEnabled); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &rp.DecryptionEnabled); err != nil {
		return err
	}
	if rp.GameDateTime, err = core.DeserializeGameDateTime(buffer); err != nil {
		return err
	}
	if err = helpers.ReadBytesAndSize(buffer, &rp.ResourceKeyArrayZipped); err != nil {
		return err
	}
	// if unzipped, err := helpers.DecompressZlib(rp.ResourceKeyArrayZipped); err != nil {
	// 	return err
	// } else {
	// 	fmt.Printf("ResourceKeyArrayZipped: %s\n", unzipped)
	// }
	if err = helpers.ReadBytesAndSize(buffer, &rp.CustomizationArrayZipped); err != nil {
		return err
	}
	// if unzipped, err := helpers.DecompressZlib(rp.CustomizationArrayZipped); err != nil {
	// 	return err
	// } else {
	// 	fmt.Printf("CustomizationArrayZipped: %s\n", unzipped)
	// }
	if err = helpers.ReadBytesAndSize(buffer, &rp.WeatherArrayZipped); err != nil {
		return err
	}
	// if unzipped, err := helpers.DecompressZlib(rp.WeatherArrayZipped); err != nil {
	// 	return err
	// } else {
	// 	fmt.Printf("WeatherArrayZipped: %s\n", unzipped)
	// }
	if err = helpers.ReadByte(buffer, &rp.Season); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &rp.CanRestart); err != nil {
		return err
	}
	if err = helpers.ReadInt32(buffer, (*int32)(&rp.EmemberCategory)); err != nil {
		return err
	}
	if err = helpers.ReadFloat32(buffer, &rp.FixedDeltaTime); err != nil {
		return err
	}
	if err = helpers.ReadBytesAndSize(buffer, &rp.InteractablesZipped); err != nil {
		return err
	}
	// if unzipped, err := helpers.DecompressZlib(rp.InteractablesZipped); err != nil {
	// 	return err
	// } else {
	// 	fmt.Printf("InteractablesZipped: %s\n", unzipped)
	// }
	if err = helpers.ReadBytesAndSize(buffer, &rp.SessionId); err != nil {
		return err
	}
	if rp.Bounds.Min, err = core.DeserializeVector3(buffer); err != nil {
		return err
	}
	if rp.Bounds.Max, err = core.DeserializeVector3(buffer); err != nil {
		return err
	}
	if err = helpers.ReadUInt16(buffer, &rp.SCPort); err != nil {
		return err
	}
	if err = helpers.ReadByte(buffer, (*byte)(&rp.EnetLogsLevel)); err != nil {
		return err
	}
	if rp.GitVersion, err = core.DeserializeGitVersion(buffer); err != nil {
		return err
	}
	if err = helpers.ReadBool(buffer, &rp.SpeedLimitEnabled); err != nil {
		return err
	}
	if rp.SpeedLimitEnabled {
		if err = rp.Config.Deserialize(buffer); err != nil {
			return err
		}
	}
	var hasVoipSettings bool
	if err = helpers.ReadBool(buffer, &hasVoipSettings); err != nil {
		return err
	}
	if hasVoipSettings {
		rp.VoipSettings = &core.VoipSettings{}
		if err = rp.VoipSettings.Deserialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (rp *PacketConnection) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteBool(buffer, rp.EncryptionEnabled); err != nil {
		return err
	}
	if err = helpers.WriteBool(buffer, rp.DecryptionEnabled); err != nil {
		return err
	}
	if err = rp.GameDateTime.Serialize(buffer); err != nil {
		return err
	}
	if err = helpers.WriteBytesAndSize(buffer, rp.ResourceKeyArrayZipped); err != nil {
		return err
	}
	if err = helpers.WriteBytesAndSize(buffer, rp.CustomizationArrayZipped); err != nil {
		return err
	}
	if err = helpers.WriteBytesAndSize(buffer, rp.WeatherArrayZipped); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, rp.Season); err != nil {
		return err
	}
	if err = helpers.WriteBool(buffer, rp.CanRestart); err != nil {
		return err
	}
	if err = helpers.WriteInt32(buffer, int32(rp.EmemberCategory)); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, rp.FixedDeltaTime); err != nil {
		return err
	}
	if err = helpers.WriteBytesAndSize(buffer, rp.InteractablesZipped); err != nil {
		return err
	}
	if err = helpers.WriteBytesAndSize(buffer, rp.SessionId); err != nil {
		return err
	}
	if err = rp.Bounds.Min.Serialize(buffer); err != nil {
		return err
	}
	if err = rp.Bounds.Max.Serialize(buffer); err != nil {
		return err
	}
	if err = helpers.WriteUInt16(buffer, rp.SCPort); err != nil {
		return err
	}
	if err = helpers.WriteByte(buffer, byte(rp.EnetLogsLevel)); err != nil {
		return err
	}
	if err = rp.GitVersion.Serialize(buffer); err != nil {
		return err
	}
	if err = helpers.WriteBool(buffer, rp.SpeedLimitEnabled); err != nil {
		return err
	}
	if rp.SpeedLimitEnabled {
		if err = rp.Config.Serialize(buffer); err != nil {
			return err
		}
	}
	hasVoipSettings := rp.VoipSettings != nil
	if err = helpers.WriteBool(buffer, hasVoipSettings); err != nil {
		return err
	}
	if hasVoipSettings {
		if err = rp.VoipSettings.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

func (rp *PacketConnection) DebugPrintAll() {
	log.Printf("EncryptionEnabled: %v\n", rp.EncryptionEnabled)
	log.Printf("DecryptionEnabled: %v\n", rp.DecryptionEnabled)
	log.Printf("GameDateTime: %v\n", rp.GameDateTime)
	// log.Printf("ResourceKeyArrayZipped: %v\n", rp.ResourceKeyArrayZipped)
	// log.Printf("CustomizationArrayZipped: %v\n", rp.CustomizationArrayZipped)
	// log.Printf("WeatherArrayZipped: %v\n", rp.WeatherArrayZipped)
	log.Printf("Season: %v\n", rp.Season)
	log.Printf("CanRestart: %v\n", rp.CanRestart)
	log.Printf("EmemberCategory: %v\n", rp.EmemberCategory)
	log.Printf("FixedDeltaTime: %v\n", rp.FixedDeltaTime)
	// log.Printf("InteractablesZipped: %v\n", rp.InteractablesZipped)
	log.Printf("SessionId: %v\n", rp.SessionId)
	log.Printf("Bounds: %v\n", rp.Bounds)
	log.Printf("SCPort: %v\n", rp.SCPort)
	log.Printf("EnetLogsLevel: %v\n", rp.EnetLogsLevel)
	log.Printf("GitVersion: %v\n", rp.GitVersion)
	log.Printf("SpeedLimitEnabled: %v\n", rp.SpeedLimitEnabled)
	log.Printf("Config: %v\n", rp.Config)
	log.Printf("VoipSettings: %v\n", rp.VoipSettings)
}

func (rp *PacketConnection) GetDefault() {

	resourceArray, err := helpers.FileToZlibCompressed(filepath.Join("static/", "customs_connection_response_resource_key.json"))
	if err != nil {
		log.Printf("Error compressing zlib: %v\n", err)
	}
	customizationArray, err := helpers.FileToZlibCompressed(filepath.Join("static/", "customs_connection_response_customizations_array.json"))
	if err != nil {
		log.Printf("Error compressing zlib: %v\n", err)
	}
	weatherArray, err := helpers.FileToZlibCompressed(filepath.Join("static/", "customs_connection_response_weather.json"))
	if err != nil {
		log.Printf("Error compressing zlib: %v\n", err)
	}
	interactablesArray, err := helpers.FileToZlibCompressed(filepath.Join("static/", "customs_connection_response_interactables.json"))
	if err != nil {
		log.Printf("Error compressing zlib: %v\n", err)
	}

	dateTimeStr := "1959-12-28 00:38:42.281721048 +0000 UTC"
	layout := "2006-01-02 15:04:05.999999999 -0700 MST"
	parsedTime, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		log.Printf("Error parsing time: %v\n", err)
	}
	rp.EncryptionEnabled = false
	rp.DecryptionEnabled = false
	rp.GameDateTime = core.GameDateTime{
		GameOnly:     true,
		GameDateTime: time.Now(),
		TimeFactor:   7,
	}
	rp.ResourceKeyArrayZipped = resourceArray
	rp.CustomizationArrayZipped = customizationArray
	rp.WeatherArrayZipped = weatherArray
	rp.Season = 3
	rp.CanRestart = false
	rp.EmemberCategory = 2
	rp.FixedDeltaTime = 0.02
	rp.InteractablesZipped = interactablesArray
	rp.SessionId = make([]byte, 256)
	rp.Bounds = core.Bounds{
		Min: core.Vector3{
			X: -448,
			Y: -250,
			Z: -280,
		},
		Max: core.Vector3{
			X: 752,
			Y: 250,
			Z: 260,
		},
	}
	rp.SCPort = 0
	rp.EnetLogsLevel = 0
	rp.GitVersion = core.GitVersion{
		CommitHash:    "6725c8dfb824531a8388fc4b88696006d0ad0132",
		CommitDate:    parsedTime,
		CommitSubject: "Export to server",
		CommitBranch:  "release/0.14.9.1",
	}
	rp.SpeedLimitEnabled = true
	rp.Config = core.ConnectionConfig{
		DefaultPlayerStateLimits: core.PlayerStateLimits{
			MinSpeed: 3.6,
			MaxSpeed: 4.6,
		},
		DictPlayerStateLimits: map[core.EPlayerState]core.PlayerStateLimits{
			0: {
				MinSpeed: 0.5,
				MaxSpeed: 1.5,
			},
			1: {
				MinSpeed: 0.5,
				MaxSpeed: 1.5,
			},
			2: {
				MinSpeed: 0.5,
				MaxSpeed: 1.5,
			},
			3: {
				MinSpeed: 1,
				MaxSpeed: 1.5,
			},
			4: {
				MinSpeed: 3.6,
				MaxSpeed: 4.6,
			},
			5: {
				MinSpeed: 6.6,
				MaxSpeed: 7.9,
			},
			6: {
				MinSpeed: 6.6,
				MaxSpeed: 7.9,
			},
			7: {
				MinSpeed: 6.6,
				MaxSpeed: 7.9,
			},
			8: {
				MinSpeed: 6.6,
				MaxSpeed: 7.9,
			},
			9: {
				MinSpeed: 0.5,
				MaxSpeed: 1.5,
			},
			10: {
				MinSpeed: 0.5,
				MaxSpeed: 1.5,
			},
			11: {
				MinSpeed: 0.5,
				MaxSpeed: 1.5,
			},
			15: {
				MinSpeed: 3.6,
				MaxSpeed: 4.6,
			},
			20: {
				MinSpeed: 0.5,
				MaxSpeed: 1.5,
			},
			22: {
				MinSpeed: 0.5,
				MaxSpeed: 1.5,
			},
		},
	}
	rp.VoipSettings = &core.VoipSettings{
		VoipEnabled: false,
		// Disabling to avoid having to deal with Dissonance HLAPI
		// VoipEnabled: true,
		VoipQualitySettings: core.VoipQualitySettings{
			FrameSize:              1,
			AudioQuality:           1,
			ForwardErrorCorrection: true,
			NoiseSuppression:       2,
			SensitivityLevels:      1,
		},
		PushToTalkSettings: core.PushToTalkSettings{
			SpeakingSecondsLimit:    20,
			SpeakingSecondsInterval: 23,
			ActivationsLimit:        4,
			ActivationsInterval:     2,
			BlockingTime:            10,
			AlertDistanceMeters:     5,
			HearingDistance:         50,
			AbuseTraceSeconds:       10,
		},
	}
}
