package config

import (
	"github.com/gin-gonic/gin"
)

type SessionData struct {
	AccountID string
	ProfileID string
}

var sessions = make(map[string]SessionData)

func generateSessionID() string {
	return "sh3-11111111111111111111111111111111"
}

func GetSessionID(c *gin.Context) string {
	cookie, err := c.Cookie("PHPSESSID")
	if err != nil {
		return ""
	}
	return cookie
}

func createSession(accountID string) {
	sessionID := generateSessionID()
	sessions[sessionID] = SessionData{AccountID: accountID, ProfileID: ""}
}

func UpdateSessionProfileID(sessionID string, profileID string) bool {
	sessionData, exists := sessions[sessionID]
	if !exists {
		return false
	}
	sessionData.ProfileID = profileID
	sessions[sessionID] = sessionData
	return true
}

func GetSession(c *gin.Context) *SessionData {
	sessionID := GetSessionID(c)
	if sessionID == "" {
		return nil
	}
	sessionData, exists := sessions[sessionID]
	if !exists {
		return nil
	}
	return &sessionData
}
