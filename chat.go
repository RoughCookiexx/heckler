package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const apiURL = "https://api.openai.com/v1/chat/completions"

type chatRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatResponse struct {
	Choices []struct {
		Message message `json:"message"`
	} `json:"choices"`
}

func sendMessageToChatGPT(systemMessageContent string, userMessageContent string) string {
	apiKey := os.Getenv("CHAT_GPT_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: CHAT_GPT_API_KEY environment variable not set.")
		os.Exit(1)
	}

	requestData := chatRequest{
		Model: "o4-mini-2025-04-16",
		Messages: []message{
			{Role: "system", Content: systemMessageContent},
			{Role: "user", Content: userMessageContent},
		},
	}

	jsonData, _ := json.Marshal(requestData) // Ignoring error as requested

	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData)) // Ignoring error
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Minimal error handling for the HTTP request itself, as failure here is common.
		return fmt.Sprintf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body) // Ignoring error

	var apiResp chatResponse
	_ = json.Unmarshal(body, &apiResp) // Ignoring error

	if len(apiResp.Choices) > 0 {
		return apiResp.Choices[0].Message.Content
	}
	return "No response content received." // Or handle as empty string
}
