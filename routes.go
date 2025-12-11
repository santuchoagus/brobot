package main

import (
	"github.com/go-telegram/bot"
	"github.com/santuchoagus/brobot/handlers"
)

var prefix = "bro"

func RegisterRoutes(b *bot.Bot) {
	// TODO: Esto es un asco pero funciona por ahora, habria que guardarlo en una DB o un json.
	b.RegisterHandler(bot.HandlerTypeMessageText, "/plan", bot.MatchTypePrefix, handlers.HandlePlan)
	b.RegisterHandler(bot.HandlerTypeMessageText, prefix+" plan", bot.MatchTypePrefix, handlers.HandlePlan)
	b.RegisterHandler(bot.HandlerTypeMessageText, prefix+" flan", bot.MatchTypePrefix, handlers.HandleFlan)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/links", bot.MatchTypePrefix, handlers.HandleLinks)
	b.RegisterHandler(bot.HandlerTypeMessageText, prefix+" links", bot.MatchTypePrefix, handlers.HandleLinks)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/contribuir", bot.MatchTypePrefix, handlers.HandleContribuir)
	b.RegisterHandler(bot.HandlerTypeMessageText, prefix+" contribuir", bot.MatchTypePrefix, handlers.HandleContribuir)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/materias", bot.MatchTypePrefix, handlers.HandleMaterias)
}
