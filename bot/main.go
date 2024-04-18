package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Token string

func init() {
	flag.StringVar(&Token, "t", "", "The discord api bot token for your bot")
	flag.Parse()
}

func main() {

	log.Info().Msg("Starting UBOT. (c) 2024 Nikolas Heise. https://sonefall.com/")
	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal().Msg("Fatal error creating discord session, exiting...: " + err.Error())
		return
	}

	log.Info().Msg("Discord session initialized. Adding Handlers...")

	bot.AddHandler(messageCreate)

	err = bot.Open()
	if err != nil {
		log.Fatal().Msg("Fatal error opening connection, exiting..., " + err.Error())
		return
	}

	log.Info().Msg("UBOT is running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
}

func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}
	if msg.Content == "/hello" {
		session.ChannelMessageSend(msg.ChannelID, "Hi!")
	}
}
