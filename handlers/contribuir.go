package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleContribuir(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	disablePreview := true

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:             update.Message.Chat.ID,
		Text:               "Colaborar con un PR o issue es bienvenido \n 👉 [Repositorio de GitHub](https://github.com/santuchoagus/brobot)",
		ParseMode:          "MarkdownV2",
		LinkPreviewOptions: &models.LinkPreviewOptions{IsDisabled: &disablePreview},
	})
}
