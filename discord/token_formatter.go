package discord

import (
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/miromax42/discord-mass-DM-GO/utilities"
)

func LaunchTokenFormatter() {
	color.Cyan("Email:Password:Token to Token")
	Tokens, err := utilities.ReadLines("tokens.txt")
	if err != nil {
		color.Red("Error while opening tokens.txt: %v", err)
		utilities.ExitSafely()
		return
	}
	if len(Tokens) == 0 {
		color.Red("[%v] Enter your tokens in tokens.txt", time.Now().Format("15:04:05"))
		utilities.ExitSafely()
		return
	}
	var onlytokens []string
	for i := 0; i < len(Tokens); i++ {
		if strings.Contains(Tokens[i], ":") {
			token := strings.Split(Tokens[i], ":")[2]
			onlytokens = append(onlytokens, token)
		}
	}
	t := utilities.TruncateLines("tokens.txt", onlytokens)
	if t != nil {
		color.Red("[%v] Error while truncating tokens.txt: %v", time.Now().Format("15:04:05"), t)
		utilities.ExitSafely()
		return
	}
}
