package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	LinkCommand = &discordgo.ApplicationCommand{
		Name:        "link",
		Description: "Link your Discord account to Esefex",
	}
)

func (c *CommandHandlers) Link(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	linkToken, err := c.dbs.LinkTokenStore.CreateToken(i.Member.User.ID)
	if err != nil {
		return nil, err
	}

	channel, err := s.UserChannelCreate(i.Member.User.ID)
	if err != nil {
		return nil, err
	}

	// https://esefex.com/link?<linktoken>
	linkUrl := fmt.Sprintf("%s/link?t=%s", c.domain, linkToken.Token)
	expiresIn := linkToken.Expiry.Sub(time.Now())
	_, err = s.ChannelMessageSend(channel.ID, fmt.Sprintf("Click this link to link your Discord account to Esefex (expires in %d Minutes): \n%s", int(expiresIn.Minutes()), linkUrl))

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Please check your DMs for a link to link your Discord account to Esefex",
		},
	}, nil

}
