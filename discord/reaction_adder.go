package discord

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/miromax42/discord-mass-DM-GO/instance"
	"github.com/miromax42/discord-mass-DM-GO/utilities"
)

func LaunchReactionAdder() {
	color.Cyan("Reaction Adder")
	color.White("Note: You don't need to do this to send DMs in servers.")
	color.White("Menu:\n1) From message\n2) Manually")
	var choice int
	fmt.Scanln(&choice)
	cfg, instances, err := instance.GetEverything()
	if err != nil {
		fmt.Println(err)
		utilities.ExitSafely()
	}
	var wg sync.WaitGroup
	wg.Add(len(instances))
	if choice == 1 {
		color.Cyan("Enter a token which can see the message:")
		var token string
		fmt.Scanln(&token)
		color.White("Enter message ID: ")
		var id string
		fmt.Scanln(&id)
		color.White("Enter channel ID: ")
		var channel string
		fmt.Scanln(&channel)
		msg, err := instance.GetRxn(channel, id, token)
		if err != nil {
			fmt.Println(err)
		}
		color.White("Select Emoji")
		for i := 0; i < len(msg.Reactions); i++ {
			color.White("%v) %v %v", i, msg.Reactions[i].Emojis.Name, msg.Reactions[i].Count)
		}
		var emoji int
		var send string
		fmt.Scanln(&emoji)
		for i := 0; i < len(instances); i++ {
			time.Sleep(time.Duration(cfg.DirectMessage.Offset) * time.Millisecond)
			go func(i int) {
				defer wg.Done()
				if msg.Reactions[emoji].Emojis.ID == "" {
					send = msg.Reactions[emoji].Emojis.Name
				} else if msg.Reactions[emoji].Emojis.ID != "" {
					send = msg.Reactions[emoji].Emojis.Name + ":" + msg.Reactions[emoji].Emojis.ID
				}
				err := instances[i].React(channel, id, send)
				if err != nil {
					fmt.Println(err)
					color.Red("[%v] %v failed to react", time.Now().Format("15:04:05"), instances[i].Token)
				} else {
					color.Green("[%v] %v reacted to the emoji", time.Now().Format("15:04:05"), instances[i].Token)
				}
			}(i)
		}
		wg.Wait()
		color.Green("[%v] Completed all threads.", time.Now().Format("15:04:05"))
	}
	if choice == 2 {
		color.Cyan("Enter channel ID")
		var channel string
		fmt.Scanln(&channel)
		color.White("Enter message ID")
		var id string
		fmt.Scanln(&id)
		color.Red("If you have a message, please use choice 1. If you want to add a custom emoji. Follow these instructions, if you don't, it won't work.\n If it's a default emoji which appears on the emoji keyboard, just copy it as TEXT not how it appears on Discord with the colons. Type it as text, it might look like 2 question marks on console but ignore.\n If it's a custom emoji (Nitro emoji) type it like this -> name:emojiID To get the emoji ID, copy the emoji link and copy the emoji ID from the URL.\nIf you do not follow this, it will not work. Don't try to do impossible things like trying to START a nitro reaction with a non-nitro account.")
		color.White("Enter emoji")
		var emoji string
		fmt.Scanln(&emoji)
		for i := 0; i < len(instances); i++ {
			time.Sleep(time.Duration(cfg.DirectMessage.Offset) * time.Millisecond)
			go func(i int) {
				defer wg.Done()
				err := instances[i].React(channel, id, emoji)
				if err != nil {
					fmt.Println(err)
					color.Red("[%v] %v failed to react", time.Now().Format("15:04:05"), instances[i].Token)
				}
				color.Green("[%v] %v reacted to the emoji", time.Now().Format("15:04:05"), instances[i].Token)
			}(i)
		}
		wg.Wait()
		color.Green("[%v] Completed all threads.", time.Now().Format("15:04:05"))
	}
}
