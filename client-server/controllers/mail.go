package controllers

import (
	"client-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MailDialogListRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type MailDialogListResponseItem struct {
	AttachmentsNew int  `json:"attachmentsNew"`
	New            int  `json:"new"`
	Type           int  `json:"type"`
	Pinned         bool `json:"pinned"`
	Message        struct {
		Dt         int    `json:"dt"`
		Type       int    `json:"type"`
		Text       string `json:"text"`
		Uid        string `json:"uid"`
		TemplateID string `json:"templateId"`
		SystemData struct {
			Date     string `json:"date"`
			Time     string `json:"time"`
			Location string `json:"location"`
		} `json:"systemData"`
	} `json:"message"`
	ID string `json:"_id"`
}

func GetMailDialogList(c *gin.Context) {
	var req MailDialogListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONResponse(c, http.StatusBadRequest, "Invalid request format", nil)
		return
	}

	messages := []MailDialogListResponseItem{}

	helpers.JSONResponse(c, http.StatusOK, "", messages)
}
