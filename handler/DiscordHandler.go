package handler

import (
	speech "DiscordSimpleChatBot/speech"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/turnage/graw/reddit"
)

// MessageCreate - Detect Message Create Event
// This function will be called (due to AddHandler) every time a new
// message is created on any channel that the autenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "/" {
		s.ChannelMessageSend(m.ChannelID, speech.IntroHelp)
	}

	if m.Content == "/momonga" {
		s.ChannelMessageSend(m.ChannelID, speech.Introduction)
	}

	if m.Content == "/meme" {
		s.ChannelMessageSend(m.ChannelID, ScrapeMemes())
	}
}

// ScrapeMemes - Scrape subreddit r/memes to retrieve latest memes
func ScrapeMemes() string {

	bot, err := reddit.NewBotFromAgentFile("./constants/discordbot.agent", 0)
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return err.Error()
	}

	harvest, err := bot.Listing("/r/memes", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/memes: ", err)
		return err.Error()
	}

	post := harvest.Posts[1:2]
	return post[0].URL
}
