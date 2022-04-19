package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
	"github.com/miromax42/discord-mass-DM-GO/discord"
)

var CaptchaServices []string

func main() {
	CaptchaServices = []string{"capmonster.cloud", "2captcha.com", "rucaptcha.com", "anti-captcha.com"}
	rand.Seed(time.Now().UTC().UnixNano())
	Options()
}

func Options() {
	color.White("Menu:\n |- 01) Invite Joiner [Token]\n |- 02) Mass DM advertiser [Token]\n |- 03) Single DM spam [Token]\n |- 04) Reaction Adder [Token]\n |- 05) Get message [Input]\n |- 06) Email:Pass:Token to Token [Email:Password:Token]\n |- 07) Token Checker [Token]\n |- 08) Guild Leaver [Token]\n |- 09) Token Onliner [Token]\n |- 10) Scraping Menu [Input]\n |- 11) Name Changer [Email:Password:Token]\n |- 12) Profile Picture Changer [Token]\n |- 13) Token Servers Check [Token]\n |- 14) Bio Changer [Token]\n |- 15) DM on React\n |- 16) Hypesquad Changer\n |- 17) Mass token changer\n")
	color.White("\nEnter your choice: ")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	default:
		color.Red("Invalid choice!")
		Options()
	case 1:
		discord.LaunchinviteJoiner()
	case 2:
		discord.LaunchMassDM()
	case 3:
		discord.LaunchSingleDM()
	case 4:
		discord.LaunchReactionAdder()
	case 5:
		discord.LaunchGetMessage()
	case 6:
		discord.LaunchTokenFormatter()
	case 7:
		discord.LaunchTokenChecker()
	case 8:
		discord.LaunchGuildLeaver()
	case 9:
		discord.LaunchTokenOnliner()
	case 10:
		discord.LaunchScraperMenu()
	case 11:
		discord.LaunchNameChanger()
	case 12:
		discord.LaunchAvatarChanger()
	case 13:
		discord.LaunchServerChecker()
	case 14:
		discord.LaunchBioChanger()
	case 15:
		discord.LaunchDMReact()
	case 16:
		discord.LaunchHypeSquadChanger()
	case 17:
		discord.LaunchTokenChanger()
	}
	time.Sleep(1 * time.Second)
	Options()
}
