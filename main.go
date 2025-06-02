package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/RoughCookiexx/twitch_chat_subscriber"
	"github.com/RoughCookiexx/gg_sse"
)

func heckle(message string)(string) {
	log.Println("Message: ", message)
	response := sendMessageToChatGPT("Would a reasonable person interpret this as rude? Respond with only 'yes' or 'no'", message)	
	log.Println("Response: ", response)
	if strings.ToLower(response) == "yes" {
		voice_response := TextToSpeech("zxrpZKR8aSGU8OrkJzzu", message)
		SendBytes(voice_response)
	}
	return ""
}

func main() {
        fmt.Println("Subscribing to chat message stream\n")
	targetURL := "http://localhost:6969/subscribe"
	filterPattern := "PRIVMSG"
	twitch_chat_subscriber.SendRequestWithCallbackAndRegex(targetURL, heckle, filterPattern)
	http.ListenAndServe(":6970", nil)
}
