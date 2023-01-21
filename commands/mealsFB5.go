package command

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/th3ph4nt0m/mensa-discord/openmensa"
)

func MealsFB5Command(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// todays date in format "2023-01-23"
	date := time.Now().Format("2006-01-02")
	meals := openmensa.GetMeals(98, date)

	embeds := []*discordgo.MessageEmbed{}

	for _, meal := range meals {
		embeds = append(embeds, &discordgo.MessageEmbed{
			Title:  meal.Name,
			Fields: generateFields(meal),
		})
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{

			Embeds: embeds,
		},
	})
}

func generateFields(meal openmensa.Meal) []*discordgo.MessageEmbedField {
	fields := []*discordgo.MessageEmbedField{}
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Kategorie",
		Value:  meal.Category,
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:  "Preis (Studierende)",
		Value: fmt.Sprintf("%s €", strconv.FormatFloat(meal.Prices.Students, 'f', 2, 64)),
	})
	return fields
}
