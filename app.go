package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc"
	"github.com/micmonay/keybd_event"
)

var twitchusername = "<ACCOUNT NAME>"
var twitchoathkey = "oauth:<OATHKEY>"

var readme = `

████████╗██████╗ ██╗██████╗ ███████╗██╗  ██╗██████╗ 
╚══██╔══╝██╔══██╗██║██╔══██╗██╔════╝╚██╗██╔╝██╔══██╗
   ██║   ██████╔╝██║██████╔╝█████╗   ╚███╔╝ ██████╔╝
   ██║   ██╔══██╗██║██╔══██╗██╔══╝   ██╔██╗ ██╔══██╗
   ██║   ██║  ██║██║██████╔╝███████╗██╔╝ ██╗██║  ██║
   ╚═╝   ╚═╝  ╚═╝╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝

----------- Twitch Chat Camera Controller -----------
              Created by StreamDoctors © 



Game must be in focus to work!

Available commands: 
- !spectator            -> Disables/Enables spectator mode 
- !cam<number>          -> Switches to desired cam (1, 2, 3, 4, 5)
- !modonly              -> Disables/Enables mod only mode


`

// variables
var spectator = false
var modmode = false
var randomcam = false
var kb, err = keybd_event.NewKeyBonding()

func sendkey() {
	errSendKey := kb.Launching()
	if errSendKey != nil {
		panic(errSendKey)
	}
}

func randomCam() {
	for {
		time.Sleep(9 * time.Second)
		if randomcam {
			rand.Seed(time.Now().UnixNano())

			key := (1 + rand.Intn(5))

			if key == 1 {
				kb.SetKeys(keybd_event.VK_1)
				sendkey()
			} else if key == 2 {
				kb.SetKeys(keybd_event.VK_2)
				sendkey()
			} else if key == 3 {
				kb.SetKeys(keybd_event.VK_3)
				sendkey()
			} else if key == 4 {
				kb.SetKeys(keybd_event.VK_4)
				sendkey()
			} else if key == 5 {
				kb.SetKeys(keybd_event.VK_5)
				sendkey()
			}
		}
	}
}

func main() {

	channel := bufio.NewScanner(os.Stdin)

	client := twitch.NewClient(twitchusername, twitchoathkey)

	if err != nil {
		panic(err)
	}

	// For linux, it is very important wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	fmt.Print(readme)

	fmt.Print("Enter channelname: ")

	channel.Scan()

	client.Join(channel.Text())

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		if message.Message[0:1] == "!" && (message.Tags["mod"] == "1" || !modmode || strings.ToLower(message.Tags["display-name"]) == strings.ToLower(channel.Text())) {

			if len(message.Message) < 4 {
				return
			}

			command := message.Message

			if command == "!cam1" {

				kb.SetKeys(keybd_event.VK_1)
				sendkey()
				fmt.Println("Switching to Camera 1")

			} else if command == "!cam2" {

				kb.SetKeys(keybd_event.VK_2)
				sendkey()
				fmt.Println("Switching to Camera 2")

			} else if command == "!cam3" {

				kb.SetKeys(keybd_event.VK_3)
				sendkey()
				fmt.Println("Switching to Camera 3")

			} else if command == "!cam4" {

				kb.SetKeys(keybd_event.VK_4)
				sendkey()
				fmt.Println("Switching to Camera 4")

			} else if command == "!cam5" {

				kb.SetKeys(keybd_event.VK_5)
				sendkey()
				fmt.Println("Switching to Camera 5")

			} else if command == "!spectator" && (message.Tags["mod"] == "1" || strings.ToLower(message.Tags["display-name"]) == strings.ToLower(channel.Text())) {

				switch spectator {
				case true:
					kb.SetKeys(keybd_event.VK_SPACE)
					sendkey()
					spectator = false
					fmt.Println("Spectator mode OFF")
				case false:
					kb.SetKeys(keybd_event.VK_SPACE)
					sendkey()
					spectator = true
					fmt.Println("Spectator mode ON")
				}

			} else if command == "!modonly" && (message.Tags["mod"] == "1" || strings.ToLower(message.Tags["display-name"]) == strings.ToLower(channel.Text())) {

				switch modmode {
				case true:
					modmode = false
					fmt.Println("Mod only mode OFF")
				case false:
					modmode = true
					fmt.Println("Mod only mode ON")
				}

			} else if command == "!randomcam" && (message.Tags["mod"] == "1" || strings.ToLower(message.Tags["display-name"]) == strings.ToLower(channel.Text())) {

				switch randomcam {
				case true:
					randomcam = false
					fmt.Println("Random cam OFF")
				case false:
					randomcam = true
					fmt.Println("Random cam ON")
				}

			} else {
				return
			}
		}
	})

	fmt.Println("Joining channel:   " + channel.Text() + "\n-----------------------------------------\n")

	go randomCam()

	errCon := client.Connect()
	if errCon != nil {
		panic(errCon)
	}

}

// Build Windows // GOOS=windows GOARCH=386 go build -o TwitchToKeybind.exe TwitchToKeybind.go
// Build Mac     // GOOS=darwin GOARCH=amd64 go build TwitchToKeybind.go
