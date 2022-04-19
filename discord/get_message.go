package discord

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/miromax42/discord-mass-DM-GO/instance"
	"github.com/miromax42/discord-mass-DM-GO/utilities"
)

func LaunchGetMessage() {
	// Uses ?around & ?limit parameters to discord's REST API to get messages to get the exact message needed
	color.Cyan("Get Message - This will get the message from Discord which you want to send.")
	color.White("Enter your token: \n")
	var token string
	fmt.Scanln(&token)
	color.White("Enter the channelID: \n")
	var channelID string
	fmt.Scanln(&channelID)
	color.White("Enter the messageID: \n")
	var messageID string
	fmt.Scanln(&messageID)
	message, err := instance.FindMessage(channelID, messageID, token)
	if err != nil {
		color.Red("Error while finding message: %v", err)
		utilities.ExitSafely()
		return
	}
	color.Green("[%v] Message: %v", time.Now().Format("15:04:05"), message)
}
