package routes

import (
	"bytes"
	"client-server/config"
	"client-server/controllers"
	"client-server/controllers/friend"
	"client-server/controllers/game"
	"client-server/controllers/hideout"
	"client-server/controllers/match"
	"client-server/controllers/menu"
	"client-server/controllers/trading"
	"compress/zlib"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// BufferMiddleware handles reassembling the entire HTTP payload and decompressing if needed
func BufferMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			return
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			contentLength := c.Request.Header.Get("Content-Length")
			requestLength, err := strconv.Atoi(contentLength)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content length"})
				c.Abort()
				return
			}

			buffer := bytes.NewBuffer(make([]byte, 0, requestLength))
			written := 0

			c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(requestLength))

			for {
				chunk := make([]byte, 1024)
				n, err := c.Request.Body.Read(chunk)
				if err != nil && err != io.EOF {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
					c.Abort()
					return
				}
				if n == 0 {
					break
				}
				buffer.Write(chunk[:n])
				written += n
			}

			requestIsCompressed := c.Request.Header.Get("requestcompressed") != "0"
			requestCompressed := c.Request.Method == "PUT" || requestIsCompressed

			var body []byte
			if requestCompressed {
				reader, err := zlib.NewReader(buffer)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create zlib reader", "details": err.Error()})
					c.Abort()
					return
				}
				defer reader.Close()

				body, err = io.ReadAll(reader)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read zlib body", "details": err.Error()})
					c.Abort()
					return
				}
			} else {
				body = buffer.Bytes()
			}

			// Log the body if not compressed
			// if !requestIsCompressed {
			// 	c.Logger().Debugf("Request body: %s", string(body))
			// }

			// Replace the request body with the decompressed body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

			// Pass the decompressed body to the next handler
			c.Set("body", body)
		}

		c.Next()
	}
}

// RequestLogger logs the request details including headers and body
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Log request headers
		fmt.Printf("Request Headers: %v\n", c.Request.Header)

		// Log request body
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Printf("Failed to read request body: %v\n", err)
		}
		// Restore the request body so other handlers can read it
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		fmt.Printf("Request Body: %s\n", string(bodyBytes))

		c.Next()

		latency := time.Since(t)

		fmt.Printf("%s %s %s %s\n",
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			latency,
		)
	}
}

func SessionValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := config.GetSessionID(c)
		if sessionID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No session found"})
			c.Abort()
			return
		}

		session := config.GetSession(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}

		// Session is valid, proceed with the request
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(BufferMiddleware())
	r.Use(RequestLogger())
	r.Use(SessionValidationMiddleware())

	r.POST("/profiles/:accountID", controllers.CreateProfile)

	r.POST("/client/languages", controllers.GetLanguages)
	r.POST("/client/items", controllers.GetItems)
	r.POST("/client/items/prices", controllers.GetItemPrices)
	r.POST("/client/customization", controllers.GetCustomization)
	r.POST("/client/globals", controllers.GetGlobals)
	r.POST("/client/settings", controllers.GetSettings)
	r.POST("/client/locale/en", controllers.GetLocaleEn)
	r.POST("/client/weather", controllers.GetWeather)
	r.POST("/client/quest/list", controllers.GetQuestList)
	r.POST("/client/locations", controllers.GetLocations)
	r.POST("/client/achievement/statistic", controllers.GetAchievementStatistic)
	r.POST("/client/achievement/list", controllers.GetAchievementList)

	r.POST("/client/menu/locale/en", menu.GetMenuLocaleEn)

	r.POST("/client/game/mode", game.GetGameMode)
	r.POST("/client/game/start", game.GameStart)
	r.POST("/client/game/version/validate", game.VersionValidate)
	r.POST("/client/game/config", game.GetGameConfig)
	r.POST("/client/game/profile/list", game.ProfileList)
	r.POST("/client/game/profile/select", game.ProfileSelect)
	r.POST("/client/handbook/templates", controllers.GetHandbookTemplates)
	r.POST("/client/builds/list", controllers.GetHandbookBuildsMyList)
	r.POST("/client/server/list", controllers.GetServerList)

	r.POST("/client/profile/status", controllers.GetProfileStatus)

	r.POST("/client/trading/api/traderSettings", trading.GetTraderSettings)
	r.POST("/client/trading/customization/storage", trading.GetCustomizationStorage)

	r.POST("/client/hideout/areas", hideout.GetHideoutAreas)
	r.POST("/client/hideout/qte/list", hideout.GetHideoutQteList)
	r.POST("/client/hideout/settings", hideout.GetHideoutSettings)
	r.POST("/client/hideout/production/recipes", hideout.GetHideoutProductionRecipes)
	r.POST("/client/hideout/production/scavcase/recipes", hideout.GetHideoutProductionScavcaseRecipes)

	r.POST("/client/repeatalbeQuests/activityPeriods", controllers.GetRepeatableQuestsActivityPeriods)

	r.POST("/client/notifier/channel/create", controllers.CreateNotifierChannel)

	r.POST("/client/friend/list", friend.GetFriendList)
	r.POST("/client/friend/request/list/inbox", friend.GetFriendRequestListInbox)
	r.POST("/client/friend/request/list/outbox", friend.GetFriendRequestListOutbox)

	r.POST("/client/mail/dialog/list", controllers.GetMailDialogList)

	r.POST("/client/match/group/current", match.GetMatchGroupCurrent)

	return r
}
