package main

import (
	"disc_ai/bot"
	"log"
	"os"
)

func main() {
	discord_token, err := os.ReadFile("discord.txt")
	if err != nil {
		log.Fatal(err)
	}
	gemini_token, err := os.ReadFile("gemini.txt")
	if err != nil {
		log.Fatal(err)
	}

	botToken := string(discord_token)
	aiKey := string(gemini_token)
	bot.BotToken = botToken
	bot.GeminiKey = aiKey
	bot.Run()
}
