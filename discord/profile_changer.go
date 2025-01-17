package discord

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/miromax42/discord-mass-DM-GO/instance"
	"github.com/miromax42/discord-mass-DM-GO/utilities"
	"github.com/zenthangplus/goccm"
)

func LaunchNameChanger() {
	color.Blue("Name Changer")
	_, instances, err := instance.GetEverything()
	if err != nil {
		color.Red("[%v] Error while getting necessary data: %v", time.Now().Format("15:04:05"), err)
	}
	for i := 0; i < len(instances); i++ {
		if instances[i].Password == "" {
			color.Red("[%v] %v No password set. It may be wrongly formatted. Only supported format is email:pass:token", time.Now().Format("15:04:05"), instances[i].Token)
			continue
		}
	}
	color.Red("NOTE: Names are changed randomly from the file.")
	users, err := utilities.ReadLines("names.txt")
	if err != nil {
		color.Red("[%v] Error while reading names.txt: %v", time.Now().Format("15:04:05"), err)
		utilities.ExitSafely()
	}
	color.Green("[%v] Enter number of threads: (0 for unlimited)", time.Now().Format("15:04:05"))

	var threads int
	fmt.Scanln(&threads)
	if threads > len(instances) || threads == 0 {
		threads = len(instances)
	}

	c := goccm.New(threads)
	for i := 0; i < len(instances); i++ {
		c.Wait()
		go func(i int) {
			err := instances[i].StartWS()
			if err != nil {
				color.Red("[%v] Error while opening websocket: %v", time.Now().Format("15:04:05"), err)
			} else {
				color.Green("[%v] Websocket opened %v", time.Now().Format("15:04:05"), instances[i].Token)
			}
			r, err := instances[i].NameChanger(users[rand.Intn(len(users))])
			if err != nil {
				color.Red("[%v] %v Error while changing name: %v", time.Now().Format("15:04:05"), instances[i].Token, err)
				return
			}
			body, err := utilities.ReadBody(r)
			if err != nil {
				fmt.Println(err)
			}
			if r.StatusCode == 200 || r.StatusCode == 204 {
				color.Green("[%v] %v Changed name successfully", time.Now().Format("15:04:05"), instances[i].Token)
			} else {
				color.Red("[%v] %v Error while changing name: %v %v", time.Now().Format("15:04:05"), instances[i].Token, r.Status, string(body))
			}
			err = instances[i].Ws.Close()
			if err != nil {
				color.Red("[%v] Error while closing websocket: %v", time.Now().Format("15:04:05"), err)
			} else {
				color.Green("[%v] Websocket closed %v", time.Now().Format("15:04:05"), instances[i].Token)
			}
			c.Done()
		}(i)
	}
	c.WaitAllDone()
	color.Green("[%v] All Done", time.Now().Format("15:04:05"))
}

func LaunchAvatarChanger() {
	color.Blue("Profile Picture Changer")
	_, instances, err := instance.GetEverything()
	if err != nil {
		color.Red("[%v] Error while getting necessary data: %v", time.Now().Format("15:04:05"), err)
	}
	color.Red("NOTE: Only PNG and JPEG/JPG supported. Profile Pictures are changed randomly from the folder. Use PNG format for faster results.")
	color.White("Loading Avatars..")
	ex, err := os.Executable()
	if err != nil {
		color.Red("Couldn't find Exe")
		utilities.ExitSafely()
	}
	ex = filepath.ToSlash(ex)
	path := path.Join(path.Dir(ex) + "/input/pfps")

	images, err := instance.GetFiles(path)
	if err != nil {
		color.Red("Couldn't find images in PFPs folder")
		utilities.ExitSafely()
	}
	color.Green("%v files found", len(images))
	var avatars []string

	for i := 0; i < len(images); i++ {
		av, err := instance.EncodeImg(images[i])
		if err != nil {
			color.Red("Couldn't encode image")
			continue
		}
		avatars = append(avatars, av)
	}
	color.Green("%v avatars loaded", len(avatars))
	color.Green("[%v] Enter number of threads: ", time.Now().Format("15:04:05"))
	var threads int
	fmt.Scanln(&threads)
	if threads > len(instances) {
		threads = len(instances)
	}

	c := goccm.New(threads)
	for i := 0; i < len(instances); i++ {
		c.Wait()
		go func(i int) {
			err := instances[i].StartWS()
			if err != nil {
				color.Red("[%v] Error while opening websocket: %v", time.Now().Format("15:04:05"), err)
			} else {
				color.Green("[%v] Websocket opened %v", time.Now().Format("15:04:05"), instances[i].Token)
			}
			r, err := instances[i].AvatarChanger(avatars[rand.Intn(len(avatars))])
			if err != nil {
				color.Red("[%v] %v Error while changing avatar: %v", time.Now().Format("15:04:05"), instances[i].Token, err)
			} else {
				if r.StatusCode == 204 || r.StatusCode == 200 {
					color.Green("[%v] %v Avatar changed successfully", time.Now().Format("15:04:05"), instances[i].Token)
				} else {
					color.Red("[%v] %v Error while changing avatar: %v", time.Now().Format("15:04:05"), instances[i].Token, r.StatusCode)
				}
			}
			err = instances[i].Ws.Close()
			if err != nil {
				color.Red("[%v] Error while closing websocket: %v", time.Now().Format("15:04:05"), err)
			} else {
				color.Green("[%v] Websocket closed %v", time.Now().Format("15:04:05"), instances[i].Token)
			}
			c.Done()
		}(i)
	}
	c.WaitAllDone()
	color.Green("[%v] All done", time.Now().Format("15:04:05"))
}

func LaunchBioChanger() {
	color.Blue("Bio changer")
	bios, err := utilities.ReadLines("bios.txt")
	if err != nil {
		color.Red("[%v] Error while reading bios.txt: %v", time.Now().Format("15:04:05"), err)
		utilities.ExitSafely()
	}
	_, instances, err := instance.GetEverything()
	if err != nil {
		color.Red("[%v] Error while getting necessary data: %v", time.Now().Format("15:04:05"), err)
		utilities.ExitSafely()
	}
	bios = instance.ValidateBios(bios)
	color.Green("[%v] Loaded %v bios, %v instances", time.Now().Format("15:04:05"), len(bios), len(instances))
	color.Green("[%v] Enter number of threads: (0 for unlimited)", time.Now().Format("15:04:05"))
	var threads int
	fmt.Scanln(&threads)
	if threads > len(instances) || threads == 0 {
		threads = len(instances)
	}
	c := goccm.New(threads)
	for i := 0; i < len(instances); i++ {
		c.Wait()
		go func(i int) {
			err := instances[i].StartWS()
			if err != nil {
				color.Red("[%v] Error while opening websocket: %v", time.Now().Format("15:04:05"), err)
			} else {
				color.Green("[%v] Websocket opened %v", time.Now().Format("15:04:05"), instances[i].Token)
			}
			err = instances[i].BioChanger(bios)
			if err != nil {
				color.Red("[%v] %v Error while changing bio: %v", time.Now().Format("15:04:05"), instances[i].Token, err)
			} else {
				color.Green("[%v] %v Bio changed successfully", time.Now().Format("15:04:05"), instances[i].Token)
			}
			err = instances[i].Ws.Close()
			if err != nil {
				color.Red("[%v] Error while closing websocket: %v", time.Now().Format("15:04:05"), err)
			} else {
				color.Green("[%v] Websocket closed %v", time.Now().Format("15:04:05"), instances[i].Token)
			}
			c.Done()
		}(i)
	}
	c.WaitAllDone()
}

func LaunchHypeSquadChanger() {
	color.Blue("Hype squad changer")
	_, instances, err := instance.GetEverything()
	if err != nil {
		color.Red("[%v] Error while getting necessary data: %v", time.Now().Format("15:04:05"), err)
		utilities.ExitSafely()
	}
	color.Green("[%v] Enter number of threads: (0 for unlimited)", time.Now().Format("15:04:05"))
	var threads int
	fmt.Scanln(&threads)
	if threads > len(instances) || threads == 0 {
		threads = len(instances)
	}
	c := goccm.New(threads)
	for i := 0; i < len(instances); i++ {
		c.Wait()
		go func(i int) {
			err := instances[i].RandomHypeSquadChanger()
			if err != nil {
				color.Red("[%v] %v Error while changing hype squad: %v", time.Now().Format("15:04:05"), instances[i].Token, err)
			} else {
				color.Green("[%v] %v Hype squad changed successfully", time.Now().Format("15:04:05"), instances[i].Token)
			}
			c.Done()
		}(i)
	}
	c.WaitAllDone()
}

func LaunchTokenChanger() {
	color.Blue("Token changer")

	_, instances, err := instance.GetEverything()
	if err != nil {
		color.Red("[%v] Error while getting necessary data: %v", time.Now().Format("15:04:05"), err)
	}
	for i := 0; i < len(instances); i++ {
		if instances[i].Password == "" {
			color.Red("[%v] %v No password set. It may be wrongly formatted. Only supported format is email:pass:token", time.Now().Format("15:04:05"), instances[i].Token)
			continue
		}
	}
	color.Green("[%v] Enter 0 to change passwords randomly and 1 to change them to a constant input", time.Now().Format("15:04:05"))
	var mode int
	fmt.Scanln(&mode)
	if mode != 0 && mode != 1 {
		color.Red("[%v] Invalid mode", time.Now().Format("15:04:05"))
		utilities.ExitSafely()
	}
	var password string
	if mode == 1 {
		color.Green("[%v] Enter Password:", time.Now().Format("15:04:05"))
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			password = scanner.Text()
		}
	}
	color.Green("[%v] Enter number of threads: (0 for unlimited)", time.Now().Format("15:04:05"))

	var threads int
	fmt.Scanln(&threads)
	if threads > len(instances) || threads == 0 {
		threads = len(instances)
	}

	c := goccm.New(threads)
	for i := 0; i < len(instances); i++ {
		c.Wait()
		go func(i int) {
			if password == "" {
				password = utilities.RandStringBytes(12)
			}
			newToken, err := instances[i].ChangeToken(password)
			if err != nil {
				color.Red("[%v] %v Error while changing token: %v", time.Now().Format("15:04:05"), instances[i].Token, err)
				err := utilities.WriteLine("input/changed_tokens.txt", fmt.Sprintf(`%s:%s:%s`, instances[i].Email, instances[i].Password, instances[i].Token))
				if err != nil {
					color.Red("[%v] Error while writing to file: %v", time.Now().Format("15:04:05"), err)
				}
			} else {
				color.Green("[%v] %v Token changed successfully", time.Now().Format("15:04:05"), instances[i].Token)
				err := utilities.WriteLine("input/changed_tokens.txt", fmt.Sprintf(`%s:%s:%s`, instances[i].Email, password, newToken))
				if err != nil {
					color.Red("[%v] Error while writing to file: %v", time.Now().Format("15:04:05"), err)
				}
			}
			c.Done()
		}(i)
	}
	c.WaitAllDone()
	color.Green("[%v] All Done", time.Now().Format("15:04:05"))
}
