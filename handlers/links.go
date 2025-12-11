package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// TODO: Hacerlo modular y que se pueda editar por afuera con archivos .md en assets, para que sea facil de contribuir.
func HandleLinks(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	disablePreview := true

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:             update.Message.Chat.ID,
		Text:               "*Links* \n• [https://taplink\\.cc/samware](https://taplink.cc/samware)\n• [Notion](https://samwareok.notion.site/)",
		ParseMode:          "MarkdownV2",
		LinkPreviewOptions: &models.LinkPreviewOptions{IsDisabled: &disablePreview},
	})
}
