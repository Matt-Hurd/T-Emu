package routes

import (
	"bytes"
	"compress/zlib"
	"eft-private-server/controllers"
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

// ResponseLogger logs the response details including headers and body
// func ResponseLogger() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Capture the response body
// 		responseBody := &bytes.Buffer{}
// 		writer := &responseBodyWriter{body: responseBody, ResponseWriter: c.Writer}
// 		c.Writer = writer

// 		c.Next()

// 		// Log response headers
// 		fmt.Printf("Response Headers: %v\n", c.Writer.Header())

// 		// Log response body
// 		fmt.Printf("Response Body: %s\n", responseBody.String())

// 		fmt.Printf("%d %s %s\n",
// 			c.Writer.Status(),
// 			c.Request.Method,
// 			c.Request.RequestURI,
// 		)
// 	}
// }

// responseBodyWriter is a custom response writer to capture the response body
// type responseBodyWriter struct {
// 	gin.ResponseWriter
// 	body *bytes.Buffer
// }

// func (w responseBodyWriter) Write(b []byte) (int, error) {
// 	w.body.Write(b)
// 	return w.ResponseWriter.Write(b)
// }

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(BufferMiddleware())
	r.Use(RequestLogger())
	// r.Use(ResponseLogger())

	r.POST("/characters/:accountID", controllers.CreateCharacter)
	r.GET("/client/game/profile/list", controllers.GetProfileCharacters)
	r.POST("/client/menu/locale/en", controllers.GetEnLocale)
	// r.POST("/users", controllers.CreateUser)

	return r
}
