package slackclient

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func SendMessage(message string) {

	token := os.Getenv("SLACK_TOKEN")
	slackChannel := os.Getenv("SLACK_CHANNEL")

	api := slack.New(token)
	attachment := slack.Attachment{
		// Pretext: message,
		Text: message,
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	channelID, timestamp, err := api.PostMessage(
		slackChannel,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	log.Printf("Message successfully sent to channel %s at %s. %s", channelID, timestamp, message)
}
