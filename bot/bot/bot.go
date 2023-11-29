package bot

import (
	"esefexbot/appcontext"
	"esefexbot/bot/actions"
	"esefexbot/msg"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Run(c *appcontext.Context) {
	log.Println("Starting bot...")

	s := c.DiscordSession

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	log.Println("Adding commands...")
	RegisterComands(s)

	go ListenForApiRequests(s, c.Channels.A2B)

	log.Println("Bot Ready.")
	<-c.Channels.Stop

	// defer actions.LeaveAllChannels(s)
	defer DeleteAllCommands(s)
}

func ListenForApiRequests(s *discordgo.Session, c chan msg.MessageA2B) {
	for {
		msg := <-c
		actions.UnprovokedMessage(s)
		log.Printf("Received message: %v", msg)
	}
}
