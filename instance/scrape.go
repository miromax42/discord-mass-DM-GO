package instance

import (
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/miromax42/discord-mass-DM-GO/utilities"
)

func Scrape(ws *Connection, Guild string, Channel string, index int) error {
	var x []interface{}
	if index == 0 {
		x = []interface{}{[2]int{0, 99}}
	} else if index == 1 {
		x = []interface{}{[2]int{0, 99}, [2]int{100, 199}}
	} else if index == 2 {
		x = []interface{}{[2]int{0, 99}, [2]int{100, 199}, [2]int{200, 299}}
	} else {
		x = []interface{}{[2]int{0, 99}, [2]int{100, 199}, [2]int{index * 100, (index * 100) + 99}}
	}

	payload := Data{
		GuildId:           Guild,
		Typing:            true,
		Threads:           true,
		Activities:        true,
		Members:           nil,
		ThreadMemberLists: nil,
		Channels: map[string]interface{}{
			Channel: x,
		},
	}

	err := ws.WriteJSONe(&Event{
		Op:   14,
		Data: payload,
	})
	if err != nil {
		return err
	}

	return nil
}

type CustomEvent struct {
	Op   int    `json:"op,omitempty"`
	Data Custom `json:"d,omitempty"`
}

type Custom struct {
	GuildID  interface{} `json:"guild_id"`
	Limit    int         `json:"limit"`
	Query    string      `json:"query"`
	Presence bool        `json:"presence"`
}

func ScrapeOffline(c *Connection, guild string, query string) error {
	custom := Custom{
		GuildID:  []string{guild},
		Limit:    100,
		Query:    query,
		Presence: true,
	}
	eventx := CustomEvent{
		Op:   8,
		Data: custom,
	}

	err := c.Conn.WriteJSON(eventx)
	if err != nil {
		return err
	}
	return nil
}

func FindNextQueries(query string, lastName string, completedQueries []string, chars string) []string {
	if query == "" {
		color.Red("[%v] Query is empty", time.Now().Format("15:04:05"))
		return nil
	}
	lastName = strings.ToLower(lastName)
	indexQuery := strings.Index(lastName, query)
	if indexQuery == -1 {
		return nil
	}
	wantedCharIndex := indexQuery + len(query)
	if wantedCharIndex >= len(lastName) {
		return nil
	}
	wantedChar := lastName[wantedCharIndex]
	queryIndexDone := strings.Index(chars, string(wantedChar))
	if queryIndexDone == -1 {
		return nil
	}

	var nextQueries []string
	for j := queryIndexDone; j < len(chars); j++ {
		newQuery := query + string(chars[j])
		if !utilities.Contains(completedQueries, newQuery) && !strings.Contains(newQuery, "  ") && string(newQuery[0]) != "" {
			nextQueries = append(nextQueries, newQuery)
		}
	}
	return nextQueries
}

func Subscribe(ws *Connection, guildid string) error {
	payload := Data{
		GuildId:    guildid,
		Typing:     true,
		Threads:    true,
		Activities: true,
	}

	err := ws.WriteJSONe(&Event{
		Op:   14,
		Data: payload,
	})
	if err != nil {
		return err
	}
	return nil
}
