package match

import (
	"bytes"
	"client-server/config"
	"client-server/helpers"
	"client-server/models"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type MatchJoinRequest struct {
	Location string `json:"location"`
	Savage   bool   `json:"savage"`
	Dt       string `json:"dt"`
	Servers  []struct {
		Ping int    `json:"ping"`
		IP   string `json:"ip"`
		Port string `json:"port"`
	} `json:"servers"`
}

type MatchJoinResponse struct {
	MaxPveCountExceeded bool                   `json:"maxPveCountExceeded"`
	Profiles            []models.ProfileStatus `json:"profiles"`
}

func GenerateProfileToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func MatchJoin(c *gin.Context) {
	var req MatchJoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONResponse(c, http.StatusBadRequest, "Invalid request format", nil)
		return
	}

	id := config.GetSession(c).ProfileID
	var profile models.Profile

	if err := config.DB.Preload("Status").First(&profile, "id = ?", id).Error; err != nil {
		helpers.JSONResponse(c, http.StatusNotFound, "Profile not found", nil)
		return
	}

	if req.Savage {
		if err := config.DB.Preload("Status").First(&profile, "id = ?", profile.SavageID).Error; err != nil {
			helpers.JSONResponse(c, http.StatusNotFound, "Profile not found", nil)
			return
		}
	}

	token := GenerateProfileToken()
	profile.Status.Status = "MatchWait"
	profile.Status.ProfileToken = &token
	profile.Status.Version = "live"
	profile.Status.Location = req.Location
	profile.Status.RaidMode = "Online"
	profile.Status.Mode = "deathmatch"
	if err := config.DB.Updates(&profile.Status).Error; err != nil {
		helpers.JSONResponse(c, http.StatusInternalServerError, "Failed to update profile status", nil)
		return
	}

	resp := MatchJoinResponse{
		MaxPveCountExceeded: false,
		Profiles: []models.ProfileStatus{
			profile.Status,
		},
	}
	helpers.JSONResponse(c, http.StatusOK, "", resp)

	//HACK
	profile.Status.IP = os.Getenv("GAME_SERVER_ADDRESS")
	port, _ := strconv.Atoi(os.Getenv("GAME_SERVER_PORT"))
	profile.Status.Port = port
	profile.Status.ServerId = os.Getenv("GAME_SERVER_ADDRESS") + "-11111_11.11.11_11.11.11"
	shortID := "ABAHBO"
	profile.Status.ShortID = &shortID
	profile.Status.AdditionalInfo = datatypes.JSON([]byte("[]"))
	profile.Status.Status = "Busy"
	if err := config.DB.Updates(&profile.Status).Error; err != nil {
		helpers.JSONResponse(c, http.StatusInternalServerError, "Failed to update profile status", nil)
		return
	}

	type JoinNotification struct {
		Type           string   `json:"type"`
		EventID        string   `json:"eventId"`
		ProfileID      string   `json:"profileid"`
		ProfileToken   string   `json:"profileToken"`
		Status         string   `json:"status"`
		IP             string   `json:"ip"`
		Port           int      `json:"port"`
		ServerID       string   `json:"sid"`
		Version        string   `json:"version"`
		Location       string   `json:"location"`
		RaidMode       string   `json:"raidMode"`
		Mode           string   `json:"mode"`
		ShortID        string   `json:"shortId"`
		AdditionalInfo []string `json:"additional_info"`
	}

	joinNotification := JoinNotification{
		Type:           "userConfirmed",
		EventID:        helpers.GenerateUUID(),
		ProfileID:      profile.ID,
		ProfileToken:   token,
		Status:         "Busy",
		IP:             profile.Status.IP,
		Port:           profile.Status.Port,
		ServerID:       profile.Status.ServerId,
		Version:        profile.Status.Version,
		Location:       profile.Status.Location,
		RaidMode:       profile.Status.RaidMode,
		Mode:           profile.Status.Mode,
		ShortID:        *profile.Status.ShortID,
		AdditionalInfo: []string{},
	}

	jsonBytes, err := json.Marshal(joinNotification)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	jsonString := string(jsonBytes)

	notifierServerURL := os.Getenv("NOTIFIER_SERVER_URL")
	url := "http://" + notifierServerURL + "/send-notification"

	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: customTransport}

	resp2, err2 := client.Post(url, "application/json", bytes.NewBuffer([]byte(
		`{"channel_id":"`+profile.NotificationChannel+`","message":"`+jsonString+`"}`,
	)))
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, err2)
		return
	}
	defer resp2.Body.Close()
}
