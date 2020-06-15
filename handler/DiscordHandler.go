package handler

import (
	speech "DiscordSimpleChatBot/speech"
	"fmt"
	"strconv"
	"strings"

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

	if strings.Fields(m.Content)[0] == "/meme" {
		s.ChannelMessageSend(m.ChannelID, ScrapeMemes(strings.Fields(m.Content)[1]))
	}
}

// ScrapeMemes - Scrape subreddit r/memes to retrieve latest memes
func ScrapeMemes(s string) string {

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

	pageNum, err := strconv.Atoi(s)
	if err != nil || pageNum < 0 {
		return "Page number is invalid (must be >0)"
	}

	post := harvest.Posts[pageNum]
	return post.URL
}
