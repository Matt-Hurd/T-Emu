package helpers

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Err    int         `json:"err"`
	ErrMsg *string     `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func JSONResponse(c *gin.Context, respCode int, errmsg string, data interface{}) {
	var errorMessage *string = &errmsg
	if respCode >= 200 && respCode < 300 {
		respCode = 0
		errorMessage = nil
	}

	response := Response{
		Err:    respCode,
		ErrMsg: errorMessage,
		Data:   data,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal response", "details": err.Error()})
		return
	}

	requestCompressed := c.Request.Method == "PUT" || c.Request.Header.Get("requestcompressed") != "0"
	if requestCompressed {
		var compressedBody bytes.Buffer
		zlibWriter := zlib.NewWriter(&compressedBody)
		if _, err := zlibWriter.Write(responseBytes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write zlib compressed response", "details": err.Error()})
			return
		}
		if err := zlibWriter.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close zlib writer", "details": err.Error()})
			return
		}

		c.Data(http.StatusOK, "application/json", compressedBody.Bytes())
	} else {
		c.Data(http.StatusOK, "application/json", responseBytes)
	}
}

func ServeJSONFile(filename string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Construct the file path
		filePath := filepath.Join("static/routes", filename)

		// Read the file
		data, err := os.ReadFile(filePath)
		if err != nil {
			JSONResponse(c, http.StatusInternalServerError, "Unable to read file", nil)
			return
		}

		// Unmarshal the JSON file content
		var jsonData interface{}
		if err := json.Unmarshal(data, &jsonData); err != nil {
			JSONResponse(c, http.StatusInternalServerError, "Failed to parse JSON file", nil)
			return
		}

		// Set content type and return the unmarshaled JSON content
		JSONResponse(c, http.StatusOK, "", jsonData)
	}
}
