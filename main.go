package main

import (
	"DiscordPfpRotator/discord"
	"DiscordPfpRotator/utils"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var once bool

	var token string
	// check if token.txt exists
	if _, err := os.Stat("token.txt"); err == nil {
		// read the token from token.txt
		tokenBytes, err := os.ReadFile("token.txt")
		if err != nil {
			log.Println("error reading token.txt: " + err.Error())
			log.Println("Press Enter to exit...")
			fmt.Scanln()
			os.Exit(1)
		}
		token = string(tokenBytes)
	} else {
		// ask for token
		fmt.Print("Enter your Discord token: ")
		fmt.Scanln(&token)
		// write the token to token.txt
		err = os.WriteFile("token.txt", []byte(token), 0644)
		if err != nil {
			log.Println("error writing token.txt: " + err.Error())
			log.Println("Press Enter to exit...")
			fmt.Scanln()
			os.Exit(1)
		}
	}

	for {
		err := filepath.Walk("images", func(path string, info fs.FileInfo, err error) error {

			// skip directories
			if info.IsDir() {
				return nil
			}

			// read the image bytes
			var fileBytes []byte
			fileBytes, err = os.ReadFile(path)
			if err != nil {
				return errors.New("error reading file: " + err.Error())
			}

			// encode the image bytes to base64
			encoded := utils.Base64Encode(string(fileBytes))

			// send the encoded image to the ChangePFP function
			err = discord.ChangePFP(token, encoded)
			if err != nil {
				return errors.New("error changing pfp: " + err.Error())
			}
			log.Println("changed pfp to " + path)

			// wait 1 minute 45 seconds
			go func() {
				for i := 0; i < 105; i++ {
					time.Sleep(1 * time.Second)
					fmt.Printf("\r%d seconds remaining", 105-i)
				}
				fmt.Println()
			}()
			time.Sleep(105 * time.Second)
			return nil
		})
		if err != nil {
			if once {
				log.Println("Something went wrong. Please check the error below and try again.")
				log.Println("If the error persists, please open an issue on GitHub.")
				log.Println("Press Enter to exit...")
				fmt.Scanln()
				os.Exit(1)
			}
			once = true
			log.Println("error walking directory: " + err.Error())
			time.Sleep(105 * time.Second)
			continue
		}
	}

}
