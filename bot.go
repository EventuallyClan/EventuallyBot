package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Ignore myself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!ev") {
		args := m.Content.Split()
		numArgs := len(args)
		//Base command
		if numArgs == 1 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "HELLO!")
			return
		}

		subCmd := args[2]

		if subCmd == "help" {
			displayHelp(s, m)
			return
		}

		if subCmd == "server" {
			subArgs := args[2:]
			if subArgs[1] == "on" {
				powerOn(subArgs[2], s, m)
			}
			if subArgs[1] == "off" {
				powerOff(subArgs[2], s, m)
			}
			return
		}
		displayHelp(s, m)
		return
	}
}

func powerOn(server string, s *discordgo.Session, m *discordgo.MessageCreate) {

}

func powerOff(server string, s *discordgo.Session, m *discordgo.MessageCreate) {

}

func displayHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "Here is where I would display a help message IF I HAD ONE")
	return
}

func main() {
	log.Print("Eventually")

	token := os.Getenv("TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("GOOOOOOOOD MORNING.  Press CTRL-C to end me.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}
