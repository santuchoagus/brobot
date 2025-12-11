package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	var (
		err      error
		b        *bot.Bot
		options  []bot.Option
		ctx      context.Context
		cancel   context.CancelFunc
		botToken string
	)

	botToken = os.Getenv("BOT_TOKEN")

	options = []bot.Option{
		bot.WithDebug(),
	}

	ctx, cancel = signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err = bot.New(botToken, options...)

	if err != nil {
		panic(err)
	}

	b.DeleteWebhook(ctx, &bot.DeleteWebhookParams{DropPendingUpdates: true})
	b.SetMyCommands(ctx, &bot.SetMyCommandsParams{Commands: Commands})
	RegisterRoutes(b)
	b.Start(ctx)
}
