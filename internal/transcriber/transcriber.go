package transcriber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	UploadUrl     = "https://api.assemblyai.com/v2/upload"
	TranscriptUrl = "https://api.assemblyai.com/v2/transcript"
)

func UploadFile(apiKey string) (string, error) {
	// Load file
	data, err := ioutil.ReadFile("test_voice_rec/test_voice.m4a")
	if err != nil {
		return "", err
	}

	// Setup HTTP client and set header
	client := &http.Client{}
	req, _ := http.NewRequest("POST", UploadUrl, bytes.NewBuffer(data))
	req.Header.Set("authorization", apiKey)
	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// Decode json and store it in a map
	var result map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	fmt.Println("success")

	// Print the upload_url
	fmt.Println(result["upload_url"])

	return fmt.Sprintf("%v", result["upload_url"]), nil
}

func Transcribe(apiKey, audioUrl string) (string, error) {
	// Prepare json data
	values := map[string]string{"audio_url": audioUrl}
	jsonData, err := json.Marshal(values)

	if err != nil {
		return "", err
	}

	// Setup HTTP client and set header
	client := &http.Client{}
	req, _ := http.NewRequest("POST", TranscriptUrl, bytes.NewBuffer(jsonData))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", apiKey)
	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// Decode json and store it in a map
	var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	fmt.Println(result["id"])
	// Print the id of the transcribed audio
	return fmt.Sprintf("%v", result["id"]), nil
}

func GetText(apiKey, transcribedId string) (string, error) {
	pollingUrl := fmt.Sprintf("%s/%s", TranscriptUrl, transcribedId)
	time.Sleep(4 * time.Second)
	// Send GET request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", pollingUrl, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", apiKey)
	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	fmt.Println(result)
	// Check status and print the transcribed text
	if result["status"] == "completed" {
		fmt.Println("success")
		fmt.Println(result["text"])
	}
	return fmt.Sprintf("%v", result["text"]), nil
}

func SetCommand() {

}
