package main

import "github.com/go-telegram/bot/models"

var Commands = []models.BotCommand{
	{
		Command:     "plan",
		Description: "Muestra un grafo del plan de estudios con sus correlativas",
	},
	{
		Command:     "materias",
		Description: "Lista los grupos de telegram por materias",
	},
	{
		Command:     "links",
		Description: "Muestra links con más información de samware",
	},
	{
		Command:     "contribuir",
		Description: "Como contribuir al desarrollo del bot",
	},
}
