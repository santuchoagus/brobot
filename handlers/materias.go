package handlers

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleMaterias(ctx context.Context, b *bot.Bot, update *models.Update) {
	type Materias struct {
		Text string
		URL  string
	}
	data, err := os.ReadFile("assets/materias.json")
	if err != nil {
		log.Println("Couldn't not seed file: assets/materias.json", err)
		return
	}

	var materias []Materias
	err = json.Unmarshal(data, &materias)
	if err != nil {
		log.Println("Error parsing json", err)
	}

	log.Println(materias)

	inlineKeyboard := [][]models.InlineKeyboardButton{}
	for i := 0; i < len(materias); i++ {
		if i%3 == 0 {
			inlineKeyboard = append(inlineKeyboard, []models.InlineKeyboardButton{})
		}
		materia := models.InlineKeyboardButton{Text: materias[i].Text, URL: materias[i].URL}
		inlineKeyboard[i/3] = append(inlineKeyboard[i/3], materia)
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		Text:   "Grupos por materias",
		ChatID: update.Message.Chat.ID,
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: inlineKeyboard,
		},
	})
}
