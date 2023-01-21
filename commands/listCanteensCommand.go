package command

import "github.com/bwmarrin/discordgo"

func ListCanteensCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{{
				Title:       "Pong!",
				Description: "Test description of a message embed using discordgo",
				Fields: []*discordgo.MessageEmbedField{
					{Name: "Field 1", Value: "Field 1 value"},
					{Name: "Field 2", Value: "Field 2 value"}}}},
		},
	})
}
