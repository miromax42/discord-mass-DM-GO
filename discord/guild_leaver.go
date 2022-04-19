package discord

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/miromax42/discord-mass-DM-GO/instance"
	"github.com/miromax42/discord-mass-DM-GO/utilities"
	"github.com/zenthangplus/goccm"
)

func LaunchGuildLeaver() {
	color.Cyan("Guild Leaver")
	cfg, instances, err := instance.GetEverything()
	if err != nil {
		color.Red("Error while getting necessary data %v", err)
		utilities.ExitSafely()

	}
	color.White("Enter the number of threads (0 for unlimited): ")
	var threads int
	fmt.Scanln(&threads)
	if threads > len(instances) {
		threads = len(instances)
	}
	if threads == 0 {
		threads = len(instances)
	}
	color.White("Enter delay between leaves: ")
	var delay int
	fmt.Scanln(&delay)
	color.White("Enter serverid: ")
	var serverid string
	fmt.Scanln(&serverid)
	c := goccm.New(threads)
	for i := 0; i < len(instances); i++ {
		time.Sleep(time.Duration(cfg.DirectMessage.Offset) * time.Millisecond)
		c.Wait()
		go func(i int) {
			p := instances[i].Leave(serverid)
			if p == 0 {
				color.Red("[%v] Error while leaving", time.Now().Format("15:04:05"))
			}
			if p == 200 || p == 204 {
				color.Green("[%v] Left server", time.Now().Format("15:04:05"))
			} else {
				color.Red("[%v] Error while leaving", time.Now().Format("15:04:05"))
			}
			time.Sleep(time.Duration(delay) * time.Second)
			c.Done()
		}(i)
	}
	c.WaitAllDone()
	color.Green("[%v] All threads finished", time.Now().Format("15:04:05"))
}
