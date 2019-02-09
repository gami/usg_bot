package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"usg_bot/command"

	"github.com/bwmarrin/discordgo"

	"usg_bot/config"
)

// Bot represents discord bot
type Bot struct {
	Token    string
	BotName  string
	commands []command.Responder
}

// NewBot is a function to initialize Bot
func NewBot(config *config.Config) *Bot {
	bot := &Bot{
		Token:    fmt.Sprintf("Bot %v", config.Token),
		BotName:  fmt.Sprintf("<@%v>", config.Token),
		commands: make([]command.Responder, 0, 8),
	}

	bot.appendCommand(&command.Hello{})
	return bot
}

// Run is a function to boot discord client
func (b *Bot) Run() error {
	client, err := discordgo.New()
	if err != nil {
		return err
	}

	client.Token = b.Token
	client.AddHandler(b.onMessage)

	err = client.Open()
	if err != nil {
		return err
	}

	fmt.Println("started bot")
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("stopped bot")

	return nil
}

func (b *Bot) appendCommand(c command.Responder) {
	b.commands = append(b.commands, c)
}

func (b *Bot) onMessage(sess *discordgo.Session, event *discordgo.MessageCreate) {
	channel, err := sess.State.Channel(event.ChannelID)
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}

	log.Println(fmt.Sprintf("received channel->%v, msg-> %v", channel.Name, event.Content))

	for _, command := range b.commands {
		if event.Author.Bot {
			continue
		}
		if command.Resolve(event.Content) {
			sendMessage(sess, channel, command.Response(event.Content))
		}
	}
}

func sendMessage(sess *discordgo.Session, channel *discordgo.Channel, msg string) {
	_, err := sess.ChannelMessageSend(channel.ID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
