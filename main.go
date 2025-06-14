package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/RoughCookiexx/twitch_chat_subscriber"
	"github.com/RoughCookiexx/gg_sse"
	"github.com/RoughCookiexx/gg_twitch_types"
)

func heckle(message twitch_types.Message)(string) {
	messageStr := message.Content
	log.Println("Message: ", messageStr)

	response := sendMessageToChatGPT("Would a reasonable person interpret this as rude? Respond with only 'yes' or 'no'", messageStr)	
	log.Println("Response: ", response)
	if strings.ToLower(response) == "yes" {
		voice_response := TextToSpeech("zxrpZKR8aSGU8OrkJzzu", messageStr)
		sse.SendBytes(voice_response)
	}
	return ""
}

func afterLastColon(s string) string {
    idx := strings.LastIndex(s, ":")
    if idx == -1 || idx+1 >= len(s) {
        return "" // or return s if you prefer to return the whole string when ':' not found
    }
    return s[idx+1:]
}

func main() {
        fmt.Println("Subscribing to chat message stream\n")
	targetURL := "http://0.0.0.0:6969/subscribe"
	filterPattern := "PRIVMSG"
	twitch_chat_subscriber.SendRequestWithCallbackAndRegex(targetURL, heckle, filterPattern, 6971)
	sse.Start()
	http.ListenAndServe((":6971"), nil)
}
