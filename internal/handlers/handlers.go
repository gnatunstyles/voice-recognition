package handlers

import (
	"fmt"
	"github.com/gnatunstyles/voice-recognition/internal/transcriber"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TranscribeHandler(c *gin.Context) {
	uploadUrl, err := transcriber.UploadFile("53c57471e8f240b29cb4cffaa9f327dc")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Error during voice file upload.",
			"error":   err,
		})
		return
	}

	transcriberId, err := transcriber.Transcribe("53c57471e8f240b29cb4cffaa9f327dc", uploadUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Error during text transcribing.",
			"error":   err,
		})
		return
	}
	fmt.Println(transcriberId)

	text, err := transcriber.GetText("53c57471e8f240b29cb4cffaa9f327dc", transcriberId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Error during getting text.",
			"error":   err,
		})
		return
	}
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Audio file transcribed succesfully.",
			"data":    gin.H{"text": text},
		})
		return
	}
}
