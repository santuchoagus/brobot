package handlers

import (
	"context"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandlePlan(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	file, err := os.Open("assets/plan.jpg")
	if err != nil {
		log.Println("Couldn't not seed file: assets/plan.jpg", err)
		return
	}

	b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID:  update.Message.Chat.ID,
		Caption: "Todavía no hay un grafo de correlativas",
		Photo: &models.InputFileUpload{
			Filename: "plan.jpg",
			Data:     file,
		},
	})
}

func HandleFlan(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	file, err := os.Open("assets/flan.png")
	if err != nil {
		log.Println("Couldn't not seed file: assets/flan.png", err)
		return
	}

	b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID:  update.Message.Chat.ID,
		Caption: "Todavía no hay un grafo de correlativas",
		Photo: &models.InputFileUpload{
			Filename: "plan.png",
			Data:     file,
		},
	})
}
