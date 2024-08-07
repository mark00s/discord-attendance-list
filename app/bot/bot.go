package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNillErr(e error) {
	//TODO: Create better error handling
	if e != nil {
		log.Fatalf("Error: %v", e)
	}
}

func Run() {
	discordClient, err := discordgo.New("Bot " + BotToken)
	checkNillErr(err)

	discordClient.AddHandler(newMessage)

	discordClient.Open()
	defer discordClient.Close()

	fmt.Println("INFO: Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discordSession *discordgo.Session, message *discordgo.MessageCreate) {
	fmt.Println("INFO: Handler")
	if message.Author.ID == discordSession.State.User.ID {
		// Don't answer myself
		return
	}

	if strings.Contains(message.Content, "!help") {
		fmt.Println("INFO: Answering !help")
		discordSession.ChannelMessageSend(message.ChannelID, "You asked for !help")
	}

	if strings.Contains(message.Content, "!bye") {
		fmt.Println("INFO: Answering !bye")
		discordSession.ChannelMessageSend(message.ChannelID, "You told me to !bye")
	}
}
